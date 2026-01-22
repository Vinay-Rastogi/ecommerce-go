package repositories

import (
	"context"
	"database/sql"

	"ecommerce/internal/models"
)

type ProductRepository interface {
	Create(ctx context.Context, product *models.ProductModel) error
	GetByID(ctx context.Context, id string) (*models.ProductModel, error)
	GetByStore(ctx context.Context, storeID string) ([]models.ProductModel, error)
	GetAll(ctx context.Context) ([]models.ProductModel, error)
}

type productRepo struct {
	db *sql.DB
}

func NewProductRepo(db *sql.DB) ProductRepository {
	return &productRepo{db}
}

// --------------------------------------
// CREATE PRODUCT
// --------------------------------------
func (r *productRepo) Create(ctx context.Context, p *models.ProductModel) error {
	query := `
		INSERT INTO products
		(id, store_id, name, description, brand, category, price, rating, availability)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		p.ID,
		p.StoreID,
		p.Name,
		p.Description,
		p.Brand,
		p.Category,
		p.Price,
		p.Rating,
		p.Availability,
	)

	return err
}

// --------------------------------------
// GET PRODUCT BY ID
// --------------------------------------
func (r *productRepo) GetByID(ctx context.Context, id string) (*models.ProductModel, error) {
	query := `
		SELECT id, store_id, name, description, brand, category,
		       price, rating, availability, created_at
		FROM products
		WHERE id = $1
	`

	var p models.ProductModel
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&p.ID,
		&p.StoreID,
		&p.Name,
		&p.Description,
		&p.Brand,
		&p.Category,
		&p.Price,
		&p.Rating,
		&p.Availability,
		&p.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &p, nil
}

// --------------------------------------
// GET PRODUCTS BY STORE
// --------------------------------------
func (r *productRepo) GetByStore(ctx context.Context, storeID string) ([]models.ProductModel, error) {
	query := `
		SELECT id, store_id, name, description, brand, category,
		       price, rating, availability, created_at
		FROM products
		WHERE store_id = $1
		ORDER BY created_at DESC
	`

	rows, err := r.db.QueryContext(ctx, query, storeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.ProductModel
	for rows.Next() {
		var p models.ProductModel
		if err := rows.Scan(
			&p.ID,
			&p.StoreID,
			&p.Name,
			&p.Description,
			&p.Brand,
			&p.Category,
			&p.Price,
			&p.Rating,
			&p.Availability,
			&p.CreatedAt,
		); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

// --------------------------------------
// GET ALL PRODUCTS (for ES indexing)
// --------------------------------------
func (r *productRepo) GetAll(ctx context.Context) ([]models.ProductModel, error) {
	query := `
		SELECT id, store_id, name, description, brand, category,
		       price, rating, availability, created_at
		FROM products
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.ProductModel
	for rows.Next() {
		var p models.ProductModel

		if err := rows.Scan(
			&p.ID,
			&p.StoreID,
			&p.Name,
			&p.Description,
			&p.Brand,
			&p.Category,
			&p.Price,
			&p.Rating,
			&p.Availability,
			&p.CreatedAt,
		); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}
