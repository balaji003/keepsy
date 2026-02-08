package users

import (
	"context"
	"errors"
)

type Service interface {
	CreateUser(ctx context.Context, req CreateUserRequest) (*User, error)
	GetUserByID(ctx context.Context, id int) (*User, error)
	CheckUser(ctx context.Context, req CheckUserRequest) (*User, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateUser(ctx context.Context, req CreateUserRequest) (*User, error) {
	if req.Name == "" || req.Email == "" {
		return nil, errors.New("name and email are required")
	}

	user := &User{
		Name:  req.Name,
		Email: req.Email,
		Phone: req.Phone,
	}

	if err := s.repo.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *service) GetUserByID(ctx context.Context, id int) (*User, error) {
	if id <= 0 {
		return nil, errors.New("invalid user ID")
	}
	return s.repo.GetByID(ctx, id)
}

func (s *service) CheckUser(ctx context.Context, req CheckUserRequest) (*User, error) {
	if req.Phone == "" {
		return nil, errors.New("phone is required")
	}
	return s.repo.GetByPhone(ctx, req.Phone)
}
