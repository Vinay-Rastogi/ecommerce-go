package repositories

import (
	"context"
	"database/sql"
	"ecommerce/internal/models"
)

type OrderRepository interface {
	Create(ctx context.Context, tx *sql.Tx, order *models.OrderModel) error
	AddItem(ctx context.Context, tx *sql.Tx, orderID string, item models.OrderItemModel) error
	GetByID(ctx context.Context, id string) (*models.OrderModel, error)

	// ✅ ADD THIS
	GetByUser(ctx context.Context, userID string) ([]models.OrderModel, error)
}


type orderRepo struct {
	db *sql.DB
}

func NewOrderRepo(db *sql.DB) OrderRepository {
	return &orderRepo{db}
}

func (r *orderRepo) Create(ctx context.Context, tx *sql.Tx, order *models.OrderModel) error {
	query := `
		INSERT INTO orders (id, user_id, status)
		VALUES ($1, $2, $3)
	`
	_, err := tx.ExecContext(ctx, query, order.ID, order.UserID, order.Status)
	return err
}

func (r *orderRepo) AddItem(
	ctx context.Context,
	tx *sql.Tx,
	orderID string,
	item models.OrderItemModel,
) error {

	query := `
		INSERT INTO order_items (id, order_id, product_id, quantity)
		VALUES (gen_random_uuid(), $1, $2, $3)
	`
	_, err := tx.ExecContext(ctx, query, orderID, item.ProductID, item.Quantity)
	return err
}

func (r *orderRepo) GetByID(ctx context.Context, id string) (*models.OrderModel, error) {
	orderQuery := `
		SELECT id, user_id, status
		FROM orders
		WHERE id = $1
	`

	var order models.OrderModel
	err := r.db.QueryRowContext(ctx, orderQuery, id).
		Scan(&order.ID, &order.UserID, &order.Status)
	if err != nil {
		return nil, err
	}

	itemsQuery := `
		SELECT product_id, quantity
		FROM order_items
		WHERE order_id = $1
	`

	rows, err := r.db.QueryContext(ctx, itemsQuery, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item models.OrderItemModel
		if err := rows.Scan(&item.ProductID, &item.Quantity); err != nil {
			return nil, err
		}
		order.Items = append(order.Items, item)
	}

	return &order, nil
}

func (r *orderRepo) GetByUser(
	ctx context.Context,
	userID string,
) ([]models.OrderModel, error) {

	query := `
		SELECT id, user_id, status
		FROM orders
		WHERE user_id = $1
		ORDER BY created_at DESC
	`

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []models.OrderModel
	for rows.Next() {
		var o models.OrderModel
		if err := rows.Scan(&o.ID, &o.UserID, &o.Status); err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}

	return orders, nil
}
