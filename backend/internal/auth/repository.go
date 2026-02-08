package auth

import (
	"context"
	"database/sql"
	"fmt"
)

type Repository interface {
	CreatePassword(ctx context.Context, userID int, hash string) error
	GetPasswordHash(ctx context.Context, userID int) (string, error)
}

type MySQLRepository struct {
	db *sql.DB
}

func NewMySQLRepository(db *sql.DB) *MySQLRepository {
	return &MySQLRepository{db: db}
}

func (r *MySQLRepository) CreatePassword(ctx context.Context, userID int, hash string) error {
	query := `INSERT INTO user_credentials (user_id, password_hash) VALUES (?, ?)`
	_, err := r.db.ExecContext(ctx, query, userID, hash)
	if err != nil {
		return fmt.Errorf("failed to create password credentials: %w", err)
	}
	return nil
}

func (r *MySQLRepository) GetPasswordHash(ctx context.Context, userID int) (string, error) {
	query := `SELECT password_hash FROM user_credentials WHERE user_id = ?`
	var hash string
	err := r.db.QueryRowContext(ctx, query, userID).Scan(&hash)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("credentials not found")
		}
		return "", fmt.Errorf("failed to get password hash: %w", err)
	}
	return hash, nil
}
