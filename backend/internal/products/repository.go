package products

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

func (r *MySQLRepository) Create(ctx context.Context, product *Product) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	query := `
		INSERT INTO products (user_id, category_id, name, brand, model, location, purchase_date, warranty_end_date, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	result, err := tx.ExecContext(ctx, query,
		product.UserID, product.CategoryID, product.Name, product.Brand, product.Model,
		product.Location, product.PurchaseDate, product.WarrantyEndDate,
		product.CreatedAt, product.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to insert product: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get last insert id: %w", err)
	}
	product.ID = int(id)

	if product.PurchaseDetails != nil {
		detailsQuery := `
			INSERT INTO product_purchase_details (product_id, shop_name, shop_address, contact_person, contact_number, order_id, delivery_status)
			VALUES (?, ?, ?, ?, ?, ?, ?)
		`
		details := product.PurchaseDetails
		_, err = tx.ExecContext(ctx, detailsQuery,
			product.ID, details.ShopName, details.ShopAddress, details.ContactPerson,
			details.ContactNumber, details.OrderID, details.DeliveryStatus,
		)
		if err != nil {
			return fmt.Errorf("failed to insert purchase details: %w", err)
		}
		// Set ProductID in the struct
		details.ProductID = product.ID
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (r *MySQLRepository) GetByID(ctx context.Context, id int) (*Product, error) {
	query := `
		SELECT id, user_id, category_id, name, brand, model, location, purchase_date, warranty_end_date, created_at, updated_at
		FROM products WHERE id = ?
	`
	var p Product
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&p.ID, &p.UserID, &p.CategoryID, &p.Name, &p.Brand, &p.Model,
		&p.Location, &p.PurchaseDate, &p.WarrantyEndDate, &p.CreatedAt, &p.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("product not found")
		}
		return nil, fmt.Errorf("failed to get product: %w", err)
	}

	// Fetch purchase details
	detailsQuery := `
		SELECT shop_name, shop_address, contact_person, contact_number, order_id, delivery_status
		FROM product_purchase_details WHERE product_id = ?
	`
	var d PurchaseDetails
	err = r.db.QueryRowContext(ctx, detailsQuery, p.ID).Scan(
		&d.ShopName, &d.ShopAddress, &d.ContactPerson, &d.ContactNumber, &d.OrderID, &d.DeliveryStatus,
	)
	if err == nil {
		d.ProductID = p.ID
		p.PurchaseDetails = &d
	} else if err != sql.ErrNoRows {
		return nil, fmt.Errorf("failed to get purchase details: %w", err)
	}

	return &p, nil
}

func (r *MySQLRepository) ListByUserID(ctx context.Context, userID int) ([]*Product, error) {
	query := `
		SELECT id, user_id, category_id, name, brand, model, location, purchase_date, warranty_end_date, created_at, updated_at
		FROM products WHERE user_id = ? ORDER BY created_at DESC
	`
	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to list products: %w", err)
	}
	defer rows.Close()

	var products []*Product
	for rows.Next() {
		var p Product
		if err := rows.Scan(
			&p.ID, &p.UserID, &p.CategoryID, &p.Name, &p.Brand, &p.Model,
			&p.Location, &p.PurchaseDate, &p.WarrantyEndDate, &p.CreatedAt, &p.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan product: %w", err)
		}
		products = append(products, &p)
	}
	return products, nil
}
