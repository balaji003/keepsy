package bills

import (
	"context"
	"database/sql"
	"fmt"
)

type Repository interface {
	Create(ctx context.Context, bill *Bill) error
	ListByUserID(ctx context.Context, userID int) ([]*Bill, error)
	GetByID(ctx context.Context, id int) (*Bill, error)
}

type mysqlRepository struct {
	db *sql.DB
}

func NewMySQLRepository(db *sql.DB) Repository {
	return &mysqlRepository{db: db}
}

func (r *mysqlRepository) Create(ctx context.Context, bill *Bill) error {
	query := `INSERT INTO keepsy_bills (product_id, file_url, file_type, created_at, updated_at) 
              VALUES (?, ?, ?, ?, ?)`

	res, err := r.db.ExecContext(ctx, query,
		bill.ProductID, bill.FileURL, bill.FileType, bill.CreatedAt, bill.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to insert bill: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get last insert id: %w", err)
	}

	bill.ID = int(id)
	return nil
}

func (r *mysqlRepository) ListByUserID(ctx context.Context, userID int) ([]*Bill, error) {
	query := `SELECT b.id, p.user_id, b.product_id, b.file_url, b.file_type, b.created_at, b.updated_at 
              FROM keepsy_bills b
              JOIN keepsy_products p ON b.product_id = p.id
              WHERE p.user_id = ? 
              ORDER BY b.created_at DESC`

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to query bills: %w", err)
	}
	defer rows.Close()

	var bills []*Bill
	for rows.Next() {
		b := &Bill{}
		if err := rows.Scan(&b.ID, &b.UserID, &b.ProductID, &b.FileURL, &b.FileType, &b.CreatedAt, &b.UpdatedAt); err != nil {
			return nil, err
		}
		bills = append(bills, b)
	}

	return bills, nil
}

func (r *mysqlRepository) GetByID(ctx context.Context, id int) (*Bill, error) {
	query := `SELECT b.id, p.user_id, b.product_id, b.file_url, b.file_type, b.created_at, b.updated_at 
              FROM keepsy_bills b
              JOIN keepsy_products p ON b.product_id = p.id
              WHERE b.id = ?`

	row := r.db.QueryRowContext(ctx, query, id)
	b := &Bill{}

	if err := row.Scan(&b.ID, &b.UserID, &b.ProductID, &b.FileURL, &b.FileType, &b.CreatedAt, &b.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("bill not found")
		}
		return nil, err
	}

	return b, nil
}
