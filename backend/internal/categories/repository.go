package categories

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

func (r *MySQLRepository) Create(ctx context.Context, category *Category) error {
	query := `
		INSERT INTO keepsy_categories (name, slug, parent_id, is_active, created_at)
		VALUES (?, ?, ?, ?, ?)
	`
	category.CreatedAt = time.Now()
	category.IsActive = true

	result, err := r.db.ExecContext(ctx, query, category.Name, category.Slug, category.ParentID, category.IsActive, category.CreatedAt)
	if err != nil {
		return fmt.Errorf("failed to create category: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get last insert id: %w", err)
	}
	category.ID = int(id)
	return nil
}

func (r *MySQLRepository) List(ctx context.Context) ([]*Category, error) {
	query := `SELECT id, name, slug, parent_id, is_active, created_at FROM keepsy_categories WHERE is_active = true ORDER BY name ASC`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to list categories: %w", err)
	}
	defer rows.Close()

	var categories []*Category
	for rows.Next() {
		var c Category
		if err := rows.Scan(&c.ID, &c.Name, &c.Slug, &c.ParentID, &c.IsActive, &c.CreatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan category: %w", err)
		}
		categories = append(categories, &c)
	}

	return categories, nil
}
