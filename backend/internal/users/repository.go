package users

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type MySQLRepository struct {
	db *sql.DB
}

func NewMySQLRepository(db *sql.DB) *MySQLRepository {
	return &MySQLRepository{db: db}
}

func (r *MySQLRepository) Create(ctx context.Context, user *User) error {
	query := `
		INSERT INTO users (uuid, name, email, phone, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	result, err := r.db.ExecContext(ctx, query, user.UUID, user.Name, user.Email, user.Phone, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get last insert id: %w", err)
	}
	user.ID = int(id)
	return nil
}

func (r *MySQLRepository) GetByID(ctx context.Context, id int) (*User, error) {
	query := `SELECT id, uuid, name, email, phone, created_at, updated_at FROM users WHERE id = ?`
	var user User
	err := r.db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.Phone, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return &user, nil
}

func (r *MySQLRepository) GetByEmail(ctx context.Context, email string) (*User, error) {
	query := `SELECT id, uuid, name, email, phone, created_at, updated_at FROM users WHERE email = ?`
	var user User
	err := r.db.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.Phone, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return &user, nil
}

func (r *MySQLRepository) GetByPhone(ctx context.Context, phone string) (*User, error) {
	query := `SELECT id, uuid, name, email, phone, created_at, updated_at FROM users WHERE phone = ?`
	var user User
	err := r.db.QueryRowContext(ctx, query, phone).Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.Phone, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return &user, nil
}

func (r *MySQLRepository) GetByUUID(ctx context.Context, uuid string) (*User, error) {
	query := `SELECT id, uuid, name, email, phone, created_at, updated_at FROM users WHERE uuid = ?`
	var user User
	err := r.db.QueryRowContext(ctx, query, uuid).Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.Phone, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return &user, nil
}
