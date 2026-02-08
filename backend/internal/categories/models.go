package categories

import (
	"context"
	"time"
)

type Category struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
	ParentID  *int      `json:"parent_id,omitempty"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateCategoryRequest struct {
	Name     string `json:"name"`
	Slug     string `json:"slug"`
	ParentID *int   `json:"parent_id,omitempty"`
}

type Repository interface {
	Create(ctx context.Context, category *Category) error
	List(ctx context.Context) ([]*Category, error)
}
