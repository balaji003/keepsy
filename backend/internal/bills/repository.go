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
	query := `INSERT INTO bills (user_id, category_id, name, file_url, file_type, amount, due_date, created_at, updated_at) 
              VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`

	res, err := r.db.ExecContext(ctx, query,
		bill.UserID, bill.CategoryID, bill.Name, bill.FileURL, bill.FileType, bill.Amount, bill.DueDate, bill.CreatedAt, bill.UpdatedAt)
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
	query := `SELECT id, user_id, category_id, name, file_url, file_type, amount, due_date, created_at, updated_at 
              FROM bills WHERE user_id = ? ORDER BY created_at DESC`

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to query bills: %w", err)
	}
	defer rows.Close()

	var bills []*Bill
	for rows.Next() {
		b := &Bill{}
		var dueDate sql.NullString   // Handle nullable date
		var categoryID sql.NullInt64 // Handle nullable category

		if err := rows.Scan(&b.ID, &b.UserID, &categoryID, &b.Name, &b.FileURL, &b.FileType, &b.Amount, &dueDate, &b.CreatedAt, &b.UpdatedAt); err != nil {
			return nil, err
		}

		if categoryID.Valid {
			catID := int(categoryID.Int64)
			b.CategoryID = &catID
		}
		if dueDate.Valid {
			dd := dueDate.String
			b.DueDate = &dd
		}

		bills = append(bills, b)
	}

	return bills, nil
}

func (r *mysqlRepository) GetByID(ctx context.Context, id int) (*Bill, error) {
	query := `SELECT id, user_id, category_id, name, file_url, file_type, amount, due_date, created_at, updated_at 
              FROM bills WHERE id = ?`

	row := r.db.QueryRowContext(ctx, query, id)
	b := &Bill{}
	var dueDate sql.NullString
	var categoryID sql.NullInt64

	if err := row.Scan(&b.ID, &b.UserID, &categoryID, &b.Name, &b.FileURL, &b.FileType, &b.Amount, &dueDate, &b.CreatedAt, &b.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("bill not found")
		}
		return nil, err
	}

	if categoryID.Valid {
		catID := int(categoryID.Int64)
		b.CategoryID = &catID
	}
	if dueDate.Valid {
		dd := dueDate.String
		b.DueDate = &dd
	}

	return b, nil
}
