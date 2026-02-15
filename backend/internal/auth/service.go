package auth

import (
	"context"
	"errors"
	"fmt"
	"keepsy-backend/internal/users"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(ctx context.Context, req RegisterRequest) (*users.User, error)
	Login(ctx context.Context, req LoginRequest) (*users.User, error)
}

type service struct {
	authRepo Repository
	userRepo users.Repository
}

func NewService(authRepo Repository, userRepo users.Repository) Service {
	return &service{
		authRepo: authRepo,
		userRepo: userRepo,
	}
}

func (s *service) Register(ctx context.Context, req RegisterRequest) (*users.User, error) {
	if req.Password == "" {
		return nil, errors.New("password is required")
	}

	// 1. Create User
	user := &users.User{
		Name:  req.Name,
		Email: req.Email,
		Phone: req.Phone,
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	// 2. Hash Password
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	// 3. Save Credentials
	if err := s.authRepo.CreatePassword(ctx, user.ID, string(hashedBytes)); err != nil {
		return nil, fmt.Errorf("failed to save credentials: %w", err)
	}

	return user, nil
}

var ErrUserNotFound = errors.New("user not found")

func (s *service) Login(ctx context.Context, req LoginRequest) (*users.User, error) {
	if req.Identifier == "" || req.Password == "" {
		return nil, errors.New("identifier and password are required")
	}

	// 1. Find User by Email OR Phone
	var user *users.User
	var err error

	// Try by Email
	user, err = s.userRepo.GetByEmail(ctx, req.Identifier)
	if err != nil {
		// Try by Phone
		user, err = s.userRepo.GetByPhone(ctx, req.Identifier)
		if err != nil {
			// If both fail, and it's because user wasn't found, return specific error
			if err.Error() == "user not found" {
				return nil, ErrUserNotFound
			}
			return nil, errors.New("invalid credentials")
		}
	}

	// 2. Get Password Hash
	hash, err := s.authRepo.GetPasswordHash(ctx, user.ID)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	// 3. Compare Password
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}
