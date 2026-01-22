package services

import (
	"context"
	"database/sql"
	"errors"
	"ecommerce/internal/models"
	"ecommerce/internal/repositories"

	"github.com/google/uuid"
)

type OrderService struct {
	db   *sql.DB
	repo repositories.OrderRepository
}

func NewOrderService(db *sql.DB,repo repositories.OrderRepository) *OrderService {
	return &OrderService{
		db:   db,
		repo: repo,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, order *models.OrderModel) error {
	if order.UserID == "" {
		return errors.New("user_id is required")
	}
	if len(order.Items) == 0 {
		return errors.New("order must contain at least one item")
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	order.ID = uuid.New().String()
	order.Status = "created"

	if err := s.repo.Create(ctx, tx, order); err != nil {
		tx.Rollback()
		return err
	}

	for _, item := range order.Items {
		if item.Quantity <= 0 {
			tx.Rollback()
			return errors.New("quantity must be greater than 0")
		}
		if err := s.repo.AddItem(ctx, tx, order.ID, item); err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func (s *OrderService) GetOrder(ctx context.Context, id string) (*models.OrderModel, error) {
	if id == "" {
		return nil, errors.New("order id is required")
	}

	if _, err := uuid.Parse(id); err != nil {
		return nil, errors.New("invalid order id format")
	}

	return s.repo.GetByID(ctx, id)
}

func (s *OrderService) GetOrdersByUser(
	ctx context.Context,
	userID string,
) ([]models.OrderModel, error) {
	return s.repo.GetByUser(ctx, userID)
}
