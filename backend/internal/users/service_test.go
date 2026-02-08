package users

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockRepo
type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) Create(ctx context.Context, user *User) error {
	args := m.Called(ctx, user)
	if args.Get(0) == nil {
		user.ID = 1
		user.CreatedAt = time.Now()
		return nil
	}
	return args.Error(0)
}

func (m *MockRepo) GetByID(ctx context.Context, id int) (*User, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*User), args.Error(1)
}

func (m *MockRepo) GetByEmail(ctx context.Context, email string) (*User, error) {
	args := m.Called(ctx, email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*User), args.Error(1)
}

func (m *MockRepo) GetByPhone(ctx context.Context, phone string) (*User, error) {
	args := m.Called(ctx, phone)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*User), args.Error(1)
}

func TestCreateUser(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockRepo := new(MockRepo)
		service := NewService(mockRepo)

		req := CreateUserRequest{
			Name:  "Test User",
			Email: "test@example.com",
			Phone: "1234567890",
		}

		mockRepo.On("Create", mock.Anything, mock.MatchedBy(func(u *User) bool {
			return u.Name == req.Name && u.Email == req.Email
		})).Return(nil)

		user, err := service.CreateUser(context.Background(), req)

		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, 1, user.ID)
		mockRepo.AssertExpectations(t)
	})

	t.Run("MissingFields", func(t *testing.T) {
		service := NewService(nil)
		_, err := service.CreateUser(context.Background(), CreateUserRequest{})
		assert.Error(t, err)
		assert.Equal(t, "name and email are required", err.Error())
	})
}

func TestGetUserByID(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockRepo := new(MockRepo)
		service := NewService(mockRepo)

		expectedUser := &User{ID: 1, Name: "Test"}
		mockRepo.On("GetByID", mock.Anything, 1).Return(expectedUser, nil)

		user, err := service.GetUserByID(context.Background(), 1)
		assert.NoError(t, err)
		assert.Equal(t, expectedUser, user)
	})

	t.Run("InvalidID", func(t *testing.T) {
		service := NewService(nil)
		_, err := service.GetUserByID(context.Background(), 0)
		assert.Error(t, err)
		assert.Equal(t, "invalid user ID", err.Error())
	})
}

func TestCheckUser(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockRepo := new(MockRepo)
		service := NewService(mockRepo)

		expectedUser := &User{ID: 1, Phone: "123"}
		mockRepo.On("GetByPhone", mock.Anything, "123").Return(expectedUser, nil)

		user, err := service.CheckUser(context.Background(), CheckUserRequest{Phone: "123"})
		assert.NoError(t, err)
		assert.Equal(t, expectedUser, user)
	})

	t.Run("MissingPhone", func(t *testing.T) {
		service := NewService(nil)
		_, err := service.CheckUser(context.Background(), CheckUserRequest{})
		assert.Error(t, err)
		assert.Equal(t, "phone is required", err.Error())
	})
}
