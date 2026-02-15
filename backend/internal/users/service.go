package users

import (
	"context"
	"errors"
)

type Service interface {
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
