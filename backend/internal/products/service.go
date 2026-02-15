package products

import (
	"context"
	"errors"
	"time"
)

type Service interface {
	CreateProduct(ctx context.Context, req CreateProductRequest) (*Product, error)
	GetProduct(ctx context.Context, id int) (*Product, error)
	ListProducts(ctx context.Context, userID int) ([]*Product, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateProduct(ctx context.Context, req CreateProductRequest) (*Product, error) {
	if req.UserID <= 0 {
		return nil, errors.New("user ID is required")
	}
	if req.Name == "" {
		return nil, errors.New("product name is required")
	}

	product := &Product{
		UserID:          req.UserID,
		CategoryID:      req.CategoryID,
		Name:            req.Name,
		Brand:           req.Brand,
		Model:           req.Model,
		Location:        req.Location,
		Price:           req.Price,
		PurchaseDate:    req.PurchaseDate,
		WarrantyEndDate: req.WarrantyEndDate,
		PurchaseDetails: req.PurchaseDetails,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	if err := s.repo.Create(ctx, product); err != nil {
		return nil, err
	}

	return product, nil
}

func (s *service) GetProduct(ctx context.Context, id int) (*Product, error) {
	if id <= 0 {
		return nil, errors.New("invalid product ID")
	}
	return s.repo.GetByID(ctx, id)
}

func (s *service) ListProducts(ctx context.Context, userID int) ([]*Product, error) {
	if userID <= 0 {
		return nil, errors.New("invalid user ID")
	}
	return s.repo.ListByUserID(ctx, userID)
}
