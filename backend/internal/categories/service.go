package categories

import (
	"context"
)

type Service interface {
	ListCategories(ctx context.Context) ([]*Category, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) ListCategories(ctx context.Context) ([]*Category, error) {
	return s.repo.List(ctx)
}
