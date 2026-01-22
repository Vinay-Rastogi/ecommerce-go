package repositories

import (
	"context"
	"database/sql"
	"ecommerce/internal/models"
)

type PaymentRepository interface {
	Create(ctx context.Context, payment *models.PaymentModel) error
}

type paymentRepo struct {
	db *sql.DB
}

func NewPaymentRepo(db *sql.DB) PaymentRepository {
	return &paymentRepo{db}
}

func (r *paymentRepo) Create(ctx context.Context, p *models.PaymentModel) error {
	query := `
		INSERT INTO payments (id, order_id, amount, status, payment_date)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := r.db.ExecContext(
		ctx,
		query,
		p.ID,
		p.OrderID,
		p.Amount,
		p.Status,
		p.PaymentDate,
	)
	return err
}
