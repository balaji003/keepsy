package bills

import (
	"context"
	"errors"
	"fmt"
	"io"
	"keepsy-backend/internal/services/storage"
	"time"
)

type Service interface {
	UploadBill(ctx context.Context, file io.Reader, filename, fileType string, req CreateBillRequest) (*Bill, error)
	ListUserBills(ctx context.Context, userID int) ([]*Bill, error)
	GetBillDownloadURL(ctx context.Context, id, userID int) (string, error)
}

type service struct {
	repo    Repository
	storage storage.Service
}

func NewService(repo Repository, storage storage.Service) Service {
	return &service{
		repo:    repo,
		storage: storage,
	}
}

func (s *service) UploadBill(ctx context.Context, file io.Reader, filename, fileType string, req CreateBillRequest) (*Bill, error) {
	// 1. Upload to Storage
	url, err := s.storage.Upload(ctx, file, filename)
	if err != nil {
		return nil, fmt.Errorf("storage upload failed: %w", err)
	}

	// 2. Create DB Record
	bill := &Bill{
		UserID:     req.UserID,
		CategoryID: req.CategoryID,
		Name:       req.Name,
		FileURL:    url,
		FileType:   fileType,
		Amount:     req.Amount,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	if err := s.repo.Create(ctx, bill); err != nil {
		// Cleanup storage if DB fails (consistency)
		_ = s.storage.Delete(ctx, url)
		return nil, fmt.Errorf("db create failed: %w", err)
	}

	return bill, nil
}

func (s *service) ListUserBills(ctx context.Context, userID int) ([]*Bill, error) {
	return s.repo.ListByUserID(ctx, userID)
}

func (s *service) GetBillDownloadURL(ctx context.Context, id, userID int) (string, error) {
	bill, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return "", err
	}

	if bill.UserID != userID {
		return "", errors.New("unauthorized access to bill")
	}

	return s.storage.GetDownloadURL(ctx, bill.FileURL)
}
