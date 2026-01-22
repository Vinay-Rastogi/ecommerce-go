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
}

type productRepo struct {
	db *sql.DB
}

func NewProductRepo(db *sql.DB) ProductRepository {
	return &productRepo{db}
}

func (r *productRepo) Create(ctx context.Context, product *models.ProductModel) error {
	query := `
		INSERT INTO products (id, store_id, name, price, availability)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := r.db.ExecContext(
		ctx,
		query,
		product.ID,
		product.StoreID,
		product.Name,
		product.Price,
		product.Availability,
	)
	return err
}

func (r *productRepo) GetByID(ctx context.Context, id string) (*models.ProductModel, error) {
	query := `
		SELECT id, store_id, name, price, availability
		FROM products
		WHERE id = $1
	`

	var product models.ProductModel
	err := r.db.QueryRowContext(ctx, query, id).
		Scan(&product.ID, &product.StoreID, &product.Name, &product.Price, &product.Availability)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *productRepo) GetByStore(
	ctx context.Context,
	storeID string,
) ([]models.ProductModel, error) {

	query := `
		SELECT id, store_id, name, price, availability
		FROM products
		WHERE store_id = $1
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
			&p.Price,
			&p.Availability,
		); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}
