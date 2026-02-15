package products

import (
	"context"
	"time"
)

type Product struct {
	ID              int              `json:"id"`
	UserID          int              `json:"user_id"`
	CategoryID      *int             `json:"category_id,omitempty"`
	Name            string           `json:"name"`
	Brand           string           `json:"brand,omitempty"`
	Model           string           `json:"model,omitempty"`
	Location        string           `json:"location,omitempty"`
	Price           *float64         `json:"price,omitempty"`
	PurchaseDate    *time.Time       `json:"purchase_date,omitempty"`
	WarrantyEndDate *time.Time       `json:"warranty_end_date,omitempty"`
	PurchaseDetails *PurchaseDetails `json:"purchase_details,omitempty"`
	CreatedAt       time.Time        `json:"created_at"`
	UpdatedAt       time.Time        `json:"updated_at"`
}

type PurchaseDetails struct {
	ProductID      int    `json:"product_id"`
	ShopName       string `json:"shop_name,omitempty"`
	ShopAddress    string `json:"shop_address,omitempty"`
	ContactPerson  string `json:"contact_person,omitempty"`
	ContactNumber  string `json:"contact_number,omitempty"`
	OrderID        string `json:"order_id,omitempty"`
	DeliveryStatus string `json:"delivery_status,omitempty"`
}

type CreateProductRequest struct {
	UserID          int              `json:"user_id"` // In real app, this comes from auth context
	CategoryID      *int             `json:"category_id,omitempty"`
	Name            string           `json:"name"`
	Brand           string           `json:"brand,omitempty"`
	Model           string           `json:"model,omitempty"`
	Location        string           `json:"location,omitempty"`
	Price           *float64         `json:"price,omitempty"`
	PurchaseDate    *time.Time       `json:"purchase_date,omitempty"`
	WarrantyEndDate *time.Time       `json:"warranty_end_date,omitempty"`
	PurchaseDetails *PurchaseDetails `json:"purchase_details,omitempty"`
}

type Repository interface {
	Create(ctx context.Context, product *Product) error
	GetByID(ctx context.Context, id int) (*Product, error)
	ListByUserID(ctx context.Context, userID int) ([]*Product, error)
}
