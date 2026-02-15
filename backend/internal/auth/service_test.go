package auth

import (
	"context"
	"errors"
	"keepsy-backend/internal/users"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

// Mocks
type MockAuthRepo struct {
	mock.Mock
}

func (m *MockAuthRepo) CreatePassword(ctx context.Context, userID int, hash string) error {
	args := m.Called(ctx, userID, hash)
	return args.Error(0)
}

func (m *MockAuthRepo) GetPasswordHash(ctx context.Context, userID int) (string, error) {
	args := m.Called(ctx, userID)
	return args.String(0), args.Error(1)
}

type MockUserRepo struct {
	mock.Mock
}

func (m *MockUserRepo) Create(ctx context.Context, user *users.User) error {
	args := m.Called(ctx, user)
	if args.Get(0) == nil {
		user.ID = 1 // Simulate ID generation
		return nil
	}
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

func TestRegister(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockAuthRepo := new(MockAuthRepo)
		mockUserRepo := new(MockUserRepo)
		service := NewService(mockAuthRepo, mockUserRepo)

		req := RegisterRequest{
			Name:     "Test User",
			Email:    "test@example.com",
			Password: "password123",
		}

		mockUserRepo.On("Create", mock.Anything, mock.MatchedBy(func(u *users.User) bool {
			return u.Name == req.Name && u.Email == req.Email
		})).Return(nil)

		mockAuthRepo.On("CreatePassword", mock.Anything, 1, mock.AnythingOfType("string")).Return(nil)

		user, err := service.Register(context.Background(), req)

		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, 1, user.ID)

		mockUserRepo.AssertExpectations(t)
		mockAuthRepo.AssertExpectations(t)
	})

	t.Run("MissingPassword", func(t *testing.T) {
		service := NewService(nil, nil)
		_, err := service.Register(context.Background(), RegisterRequest{Name: "User", Email: "e"})
		assert.Error(t, err)
		assert.Equal(t, "password is required", err.Error())
	})
}

func TestLogin(t *testing.T) {
	t.Run("SuccessEmail", func(t *testing.T) {
		mockAuthRepo := new(MockAuthRepo)
		mockUserRepo := new(MockUserRepo)
		service := NewService(mockAuthRepo, mockUserRepo)

		password := "password123"
		hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

		user := &users.User{ID: 1, Email: "test@example.com"}

		mockUserRepo.On("GetByEmail", mock.Anything, "test@example.com").Return(user, nil)
		mockAuthRepo.On("GetPasswordHash", mock.Anything, 1).Return(string(hash), nil)

		loggedInUser, err := service.Login(context.Background(), LoginRequest{
			Identifier: "test@example.com",
			Password:   password,
		})

		assert.NoError(t, err)
		assert.Equal(t, user, loggedInUser)
	})

	t.Run("WrongPassword", func(t *testing.T) {
		mockAuthRepo := new(MockAuthRepo)
		mockUserRepo := new(MockUserRepo)
		service := NewService(mockAuthRepo, mockUserRepo)

		password := "password123"
		hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

		user := &users.User{ID: 1, Email: "test@example.com"}

		mockUserRepo.On("GetByEmail", mock.Anything, "test@example.com").Return(user, nil)
		mockAuthRepo.On("GetPasswordHash", mock.Anything, 1).Return(string(hash), nil)

		_, err := service.Login(context.Background(), LoginRequest{
			Identifier: "test@example.com",
			Password:   "wrongpassword",
		})

		assert.Error(t, err)
		assert.Equal(t, "invalid credentials", err.Error())
	})

	t.Run("UserNotFound", func(t *testing.T) {
		mockAuthRepo := new(MockAuthRepo)
		mockUserRepo := new(MockUserRepo)
		service := NewService(mockAuthRepo, mockUserRepo)

		mockUserRepo.On("GetByEmail", mock.Anything, "unknown@example.com").Return(nil, errors.New("user not found"))
		mockUserRepo.On("GetByPhone", mock.Anything, "unknown@example.com").Return(nil, errors.New("user not found"))

		_, err := service.Login(context.Background(), LoginRequest{
			Identifier: "unknown@example.com",
			Password:   "password",
		})

		assert.Error(t, err)
		assert.Equal(t, ErrUserNotFound, err)
	})
}
