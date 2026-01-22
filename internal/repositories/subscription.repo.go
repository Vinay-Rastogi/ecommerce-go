package repositories

import (
	"context"
	"database/sql"
	"ecommerce/internal/models"
)

type SubscriptionRepository interface {
	Create(ctx context.Context, sub *models.SubscriptionModel) error
	GetByUser(ctx context.Context, userID string) ([]models.SubscriptionModel, error)
}

type subscriptionRepo struct {
	db *sql.DB
}

func NewSubscriptionRepo(db *sql.DB) SubscriptionRepository {
	return &subscriptionRepo{db}
}

func (r *subscriptionRepo) Create(ctx context.Context, sub *models.SubscriptionModel) error {
	query := `
		INSERT INTO subscriptions
		(id, user_id, product_id, start_date, end_date, status)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := r.db.ExecContext(
		ctx,
		query,
		sub.ID,
		sub.UserID,
		sub.ProductID,
		sub.StartDate,
		sub.EndDate,
		sub.Status,
	)
	return err
}

func (r *subscriptionRepo) GetByUser(
	ctx context.Context,
	userID string,
) ([]models.SubscriptionModel, error) {

	query := `
		SELECT id, user_id, product_id, start_date, end_date, status
		FROM subscriptions
		WHERE user_id = $1
	`

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subs []models.SubscriptionModel
	for rows.Next() {
		var s models.SubscriptionModel
		if err := rows.Scan(
			&s.ID,
			&s.UserID,
			&s.ProductID,
			&s.StartDate,
			&s.EndDate,
			&s.Status,
		); err != nil {
			return nil, err
		}
		subs = append(subs, s)
	}

	return subs, nil
}

func (r *subscriptionRepo) GetActiveByUser(
	ctx context.Context,
	userID string,
) ([]models.SubscriptionModel, error) {

	query := `
		SELECT id, user_id, product_id, start_date, end_date, status
		FROM subscriptions
		WHERE user_id = $1 AND status = 'active'
	`

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subs []models.SubscriptionModel
	for rows.Next() {
		var s models.SubscriptionModel
		if err := rows.Scan(
			&s.ID,
			&s.UserID,
			&s.ProductID,
			&s.StartDate,
			&s.EndDate,
			&s.Status,
		); err != nil {
			return nil, err
		}
		subs = append(subs, s)
	}

	return subs, nil
}

