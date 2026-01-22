package repositories

import (
	"context"
	"database/sql"
	"ecommerce/internal/models"
)

type StoreRepository interface {
	Create(ctx context.Context, store *models.StoreModel) error
	GetAll(ctx context.Context) ([]models.StoreModel, error)
}

type storeRepo struct {
	db *sql.DB
}

func NewStoreRepo(db *sql.DB) StoreRepository {
	return &storeRepo{db}
}

func (r *storeRepo) Create(ctx context.Context, store *models.StoreModel) error {
	query := `
		INSERT INTO stores (id, name, status)
		VALUES ($1, $2, $3)
	`
	_, err := r.db.ExecContext(ctx, query,
		store.ID, store.Name, store.Status,
	)
	return err
}

func (r *storeRepo) GetAll(ctx context.Context) ([]models.StoreModel, error) {
	query := `SELECT id, name, status FROM stores`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stores []models.StoreModel
	for rows.Next() {
		var store models.StoreModel
		if err := rows.Scan(&store.ID, &store.Name, &store.Status); err != nil {
			return nil, err
		}
		stores = append(stores, store)
	}

	return stores, nil
}
