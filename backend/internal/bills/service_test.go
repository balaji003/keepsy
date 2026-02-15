package bills

import (
	"context"
	"errors"
	"io"
	"strings"
	"testing"

	"keepsy-backend/internal/users"

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

// MockUserRepo
type MockUserRepo struct {
	mock.Mock
}

func (m *MockUserRepo) Create(ctx context.Context, user *users.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepo) GetByID(ctx context.Context, id int) (*users.User, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*users.User), args.Error(1)
}

func (m *MockUserRepo) GetByEmail(ctx context.Context, email string) (*users.User, error) {
	args := m.Called(ctx, email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*users.User), args.Error(1)
}

func (m *MockUserRepo) GetByPhone(ctx context.Context, phone string) (*users.User, error) {
	args := m.Called(ctx, phone)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*users.User), args.Error(1)
}

func (m *MockUserRepo) GetByUUID(ctx context.Context, uuid string) (*users.User, error) {
	args := m.Called(ctx, uuid)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*users.User), args.Error(1)
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
		mockUserRepo := new(MockUserRepo)
		mockStorage := new(MockStorage)
		service := NewService(mockRepo, mockUserRepo, mockStorage)

		fileContent := "dummy content"
		file := strings.NewReader(fileContent)
		filename := "test.pdf"
		fileType := "application/pdf"
		req := CreateBillRequest{UserID: 1, ProductID: 100}

		// Expect User Fetch
		mockUserRepo.On("GetByID", mock.Anything, 1).Return(&users.User{ID: 1, UUID: "test-uuid"}, nil)

		// Expect upload to storage with UUID path
		expectedPath := "test-uuid/bills/test.pdf"
		mockStorage.On("Upload", mock.Anything, file, expectedPath).Return("http://storage/test-uuid/bills/test.pdf", nil)

		// Expect DB creation
		mockRepo.On("Create", mock.Anything, mock.MatchedBy(func(b *Bill) bool {
			return b.FileURL == "http://storage/test-uuid/bills/test.pdf" && b.ProductID == 100
		})).Return(nil)

		bill, err := service.UploadBill(context.Background(), file, filename, fileType, req)

		assert.NoError(t, err)
		assert.NotNil(t, bill)
		assert.Equal(t, "http://storage/test-uuid/bills/test.pdf", bill.FileURL)

		mockRepo.AssertExpectations(t)
		mockUserRepo.AssertExpectations(t)
		mockStorage.AssertExpectations(t)
	})

	t.Run("StorageFailure", func(t *testing.T) {
		mockRepo := new(MockRepo)
		mockUserRepo := new(MockUserRepo)
		mockStorage := new(MockStorage)
		service := NewService(mockRepo, mockUserRepo, mockStorage)

		file := strings.NewReader("content")

		// Expect User Fetch
		mockUserRepo.On("GetByID", mock.Anything, 1).Return(&users.User{ID: 1, UUID: "test-uuid"}, nil)

		expectedPath := "test-uuid/bills/test.pdf"
		mockStorage.On("Upload", mock.Anything, file, expectedPath).Return("", errors.New("upload failed"))

		_, err := service.UploadBill(context.Background(), file, "test.pdf", "application/pdf", CreateBillRequest{UserID: 1})

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "upload failed")
		mockRepo.AssertNotCalled(t, "Create")
	})

	t.Run("DBFailure_Cleanup", func(t *testing.T) {
		mockRepo := new(MockRepo)
		mockUserRepo := new(MockUserRepo)
		mockStorage := new(MockStorage)
		service := NewService(mockRepo, mockUserRepo, mockStorage)

		file := strings.NewReader("content")

		// Expect User Fetch
		mockUserRepo.On("GetByID", mock.Anything, 1).Return(&users.User{ID: 1, UUID: "test-uuid"}, nil)

		expectedPath := "test-uuid/bills/test.pdf"
		mockStorage.On("Upload", mock.Anything, file, expectedPath).Return("http://url", nil)
		mockRepo.On("Create", mock.Anything, mock.Anything).Return(errors.New("db failed"))
		mockStorage.On("Delete", mock.Anything, "http://url").Return(nil)

		_, err := service.UploadBill(context.Background(), file, "test.pdf", "application/pdf", CreateBillRequest{UserID: 1})

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "db failed")
		mockStorage.AssertCalled(t, "Delete", mock.Anything, "http://url")
	})
}

func TestGetBillDownloadURL(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockRepo := new(MockRepo)
		mockUserRepo := new(MockUserRepo)
		mockStorage := new(MockStorage)
		service := NewService(mockRepo, mockUserRepo, mockStorage)

		bill := &Bill{ID: 1, UserID: 1, FileURL: "http://url/file.pdf"}
		mockRepo.On("GetByID", mock.Anything, 1).Return(bill, nil)
		mockStorage.On("GetDownloadURL", mock.Anything, "http://url/file.pdf").Return("http://signed-url/file.pdf", nil)

		url, err := service.GetBillDownloadURL(context.Background(), 1, 1)

		assert.NoError(t, err)
		assert.Equal(t, "http://signed-url/file.pdf", url)
	})

	t.Run("Unauthorized", func(t *testing.T) {
		mockRepo := new(MockRepo)
		mockUserRepo := new(MockUserRepo)
		mockStorage := new(MockStorage)
		service := NewService(mockRepo, mockUserRepo, mockStorage)

		bill := &Bill{ID: 1, UserID: 1, FileURL: "http://url/file.pdf"}
		mockRepo.On("GetByID", mock.Anything, 1).Return(bill, nil)

		_, err := service.GetBillDownloadURL(context.Background(), 1, 999) // Different user

		assert.Error(t, err)
		assert.Equal(t, "unauthorized access to bill", err.Error())
	})

	t.Run("NotFound", func(t *testing.T) {
		mockRepo := new(MockRepo)
		mockUserRepo := new(MockUserRepo)
		mockStorage := new(MockStorage)
		service := NewService(mockRepo, mockUserRepo, mockStorage)

		mockRepo.On("GetByID", mock.Anything, 1).Return(nil, errors.New("not found"))

		_, err := service.GetBillDownloadURL(context.Background(), 1, 1)

		assert.Error(t, err)
		assert.Equal(t, "not found", err.Error())
	})
}
