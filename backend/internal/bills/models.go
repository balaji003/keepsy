package bills

import (
	"time"
)

type Bill struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"` // Populated via JOIN
	ProductID int       `json:"product_id"`
	FileURL   string    `json:"file_url"`
	FileType  string    `json:"file_type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateBillRequest struct {
	UserID    int `json:"-"` // From Context/Auth
	ProductID int `json:"product_id"`
}
