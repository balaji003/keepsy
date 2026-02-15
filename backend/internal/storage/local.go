package storage

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

type LocalStorage struct {
	basePath string
	baseURL  string
}

// NewLocalStorage creates a new instance of LocalStorage.
// basePath: directory to save files (e.g., "./uploads")
// baseURL: base URL to access files (e.g., "http://localhost:8080/uploads")
func NewLocalStorage(basePath, baseURL string) (*LocalStorage, error) {
	// Ensure upload directory exists
	if err := os.MkdirAll(basePath, 0755); err != nil {
		return nil, fmt.Errorf("failed to create upload directory: %w", err)
	}

	return &LocalStorage{
		basePath: basePath,
		baseURL:  baseURL,
	}, nil
}

func (s *LocalStorage) Upload(ctx context.Context, file io.Reader, filename string) (string, error) {
	// Generate unique filename to avoid collisions
	// Format: timestamp_original_filename
	uniqueName := fmt.Sprintf("%d_%s", time.Now().UnixNano(), filename)
	filePath := filepath.Join(s.basePath, uniqueName)

	// Create destination file
	dst, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to create file on disk: %w", err)
	}
	defer dst.Close()

	// Copy content
	if _, err := io.Copy(dst, file); err != nil {
		return "", fmt.Errorf("failed to save file content: %w", err)
	}

	// Return public URL (relative or absolute based on baseURL)
	fileURL := fmt.Sprintf("%s/%s", s.baseURL, uniqueName)
	return fileURL, nil
}

func (s *LocalStorage) Delete(ctx context.Context, fileURL string) error {
	// Extract filename from URL (simple implementation assuming flat structure)
	// In production, might need more robust parsing logic
	filename := filepath.Base(fileURL)
	filePath := filepath.Join(s.basePath, filename)

	if err := os.Remove(filePath); err != nil {
		if os.IsNotExist(err) {
			return nil // File already gone
		}
		return fmt.Errorf("failed to delete file: %w", err)
	}

	return nil
}

func (s *LocalStorage) GetDownloadURL(ctx context.Context, fileURL string) (string, error) {
	// For local storage, the fileURL stored in DB is already the public URL
	// So we just return it.
	// If we stored relative paths, we would append baseURL here.
	return fileURL, nil
}
