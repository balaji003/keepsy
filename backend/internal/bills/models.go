package bills

import (
	"time"
)

type Bill struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	CategoryID *int      `json:"category_id,omitempty"`
	Name       string    `json:"name"`
	FileURL    string    `json:"file_url"`
	FileType   string    `json:"file_type"`
	Amount     *float64  `json:"amount,omitempty"`
	DueDate    *string   `json:"due_date,omitempty"` // simplified date handling
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type CreateBillRequest struct {
	UserID     int      `json:"-"` // From Context/Auth
	CategoryID *int     `json:"category_id"`
	Name       string   `json:"name"`
	Amount     *float64 `json:"amount"`
}
