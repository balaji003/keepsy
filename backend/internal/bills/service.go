package bills

import (
	"context"
	"errors"
	"fmt"
	"io"
	"keepsy-backend/internal/services/storage"
	"keepsy-backend/internal/users"
	"time"
)

type Service interface {
	UploadBill(ctx context.Context, file io.Reader, filename, fileType string, req CreateBillRequest) (*Bill, error)
	ListUserBills(ctx context.Context, userID int) ([]*Bill, error)
	GetBillDownloadURL(ctx context.Context, id, userID int) (string, error)
}

type service struct {
	repo     Repository
	userRepo users.Repository
	storage  storage.Service
}

func NewService(repo Repository, userRepo users.Repository, storage storage.Service) Service {
	return &service{
		repo:     repo,
		userRepo: userRepo,
		storage:  storage,
	}
}

func (s *service) UploadBill(ctx context.Context, file io.Reader, filename, fileType string, req CreateBillRequest) (*Bill, error) {
	// 0. Get User UUID
	user, err := s.userRepo.GetByID(ctx, req.UserID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	// 1. Upload to Storage (Path: <uuid>/bills/<filename>)
	storagePath := fmt.Sprintf("%s/bills/%s", user.UUID, filename)
	url, err := s.storage.Upload(ctx, file, storagePath)
	if err != nil {
		return nil, fmt.Errorf("storage upload failed: %w", err)
	}

	// 2. Create DB Record
	bill := &Bill{
		UserID:    req.UserID,
		ProductID: req.ProductID,
		FileURL:   url,
		FileType:  fileType,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
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
