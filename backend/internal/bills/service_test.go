package bills

import (
	"context"
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockRepo
type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) Create(ctx context.Context, bill *Bill) error {
	args := m.Called(ctx, bill)
	if args.Get(0) == nil {
		bill.ID = 1
		return nil
	}
	return args.Error(0)
}

func (m *MockRepo) ListByUserID(ctx context.Context, userID int) ([]*Bill, error) {
	args := m.Called(ctx, userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*Bill), args.Error(1)
}

func (m *MockRepo) GetByID(ctx context.Context, id int) (*Bill, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Bill), args.Error(1)
}

// MockStorage
type MockStorage struct {
	mock.Mock
}

func (m *MockStorage) Upload(ctx context.Context, file io.Reader, filename string) (string, error) {
	args := m.Called(ctx, file, filename)
	return args.String(0), args.Error(1)
}

func (m *MockStorage) Delete(ctx context.Context, url string) error {
	args := m.Called(ctx, url)
	return args.Error(0)
}

func (m *MockStorage) GetDownloadURL(ctx context.Context, url string) (string, error) {
	args := m.Called(ctx, url)
	return args.String(0), args.Error(1)
}

func TestUploadBill(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockRepo := new(MockRepo)
		mockStorage := new(MockStorage)
		service := NewService(mockRepo, mockStorage)

		fileContent := "dummy content"
		file := strings.NewReader(fileContent)
		filename := "test.pdf"
		fileType := "application/pdf"
		req := CreateBillRequest{UserID: 1, Name: "Bill 1"}

		// Expect upload to storage
		mockStorage.On("Upload", mock.Anything, file, filename).Return("http://storage/test.pdf", nil)

		// Expect DB creation
		mockRepo.On("Create", mock.Anything, mock.MatchedBy(func(b *Bill) bool {
			return b.FileURL == "http://storage/test.pdf" && b.Name == "Bill 1"
		})).Return(nil)

		bill, err := service.UploadBill(context.Background(), file, filename, fileType, req)

		assert.NoError(t, err)
		assert.NotNil(t, bill)
		assert.Equal(t, "http://storage/test.pdf", bill.FileURL)

		mockRepo.AssertExpectations(t)
		mockStorage.AssertExpectations(t)
	})

	t.Run("StorageFailure", func(t *testing.T) {
		mockRepo := new(MockRepo)
		mockStorage := new(MockStorage)
		service := NewService(mockRepo, mockStorage)

		file := strings.NewReader("content")
		mockStorage.On("Upload", mock.Anything, file, "test.pdf").Return("", errors.New("upload failed"))

		_, err := service.UploadBill(context.Background(), file, "test.pdf", "application/pdf", CreateBillRequest{})

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "upload failed")
		mockRepo.AssertNotCalled(t, "Create")
	})

	t.Run("DBFailure_Cleanup", func(t *testing.T) {
		mockRepo := new(MockRepo)
		mockStorage := new(MockStorage)
		service := NewService(mockRepo, mockStorage)

		file := strings.NewReader("content")
		mockStorage.On("Upload", mock.Anything, file, "test.pdf").Return("http://url", nil)
		mockRepo.On("Create", mock.Anything, mock.Anything).Return(errors.New("db failed"))
		mockStorage.On("Delete", mock.Anything, "http://url").Return(nil)

		_, err := service.UploadBill(context.Background(), file, "test.pdf", "application/pdf", CreateBillRequest{})

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "db failed")
		mockStorage.AssertCalled(t, "Delete", mock.Anything, "http://url")
	})
}

func TestGetBillDownloadURL(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockRepo := new(MockRepo)
		mockStorage := new(MockStorage)
		service := NewService(mockRepo, mockStorage)

		bill := &Bill{ID: 1, UserID: 1, FileURL: "http://url/file.pdf"}
		mockRepo.On("GetByID", mock.Anything, 1).Return(bill, nil)
		mockStorage.On("GetDownloadURL", mock.Anything, "http://url/file.pdf").Return("http://signed-url/file.pdf", nil)

		url, err := service.GetBillDownloadURL(context.Background(), 1, 1)

		assert.NoError(t, err)
		assert.Equal(t, "http://signed-url/file.pdf", url)
	})

	t.Run("Unauthorized", func(t *testing.T) {
		mockRepo := new(MockRepo)
		mockStorage := new(MockStorage)
		service := NewService(mockRepo, mockStorage)

		bill := &Bill{ID: 1, UserID: 1, FileURL: "http://url/file.pdf"}
		mockRepo.On("GetByID", mock.Anything, 1).Return(bill, nil)

		_, err := service.GetBillDownloadURL(context.Background(), 1, 999) // Different user

		assert.Error(t, err)
		assert.Equal(t, "unauthorized access to bill", err.Error())
	})

	t.Run("NotFound", func(t *testing.T) {
		mockRepo := new(MockRepo)
		mockStorage := new(MockStorage)
		service := NewService(mockRepo, mockStorage)

		mockRepo.On("GetByID", mock.Anything, 1).Return(nil, errors.New("not found"))

		_, err := service.GetBillDownloadURL(context.Background(), 1, 1)

		assert.Error(t, err)
		assert.Equal(t, "not found", err.Error())
	})
}
