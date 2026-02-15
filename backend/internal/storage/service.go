package storage

import (
	"context"
	"io"
)

// Service defines the interface for file storage operations.
// This allows switching between Local, S3, GCS, etc.
type Service interface {
	// Upload saves the file and returns the public URL/path.
	Upload(ctx context.Context, file io.Reader, filename string) (string, error)

	// Delete removes the file from storage.
	Delete(ctx context.Context, url string) error

	// GetDownloadURL returns a URL to download the file.
	// For Local: Returns the public static URL.
	// For S3: Returns a presigned URL.
	GetDownloadURL(ctx context.Context, url string) (string, error)
}
