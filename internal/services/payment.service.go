package services

import (
	"context"
	"errors"
	"time"

	"ecommerce/internal/models"
	"ecommerce/internal/repositories"

	"github.com/google/uuid"
)

type PaymentService struct {
	repo repositories.PaymentRepository
}

func NewPaymentService(repo repositories.PaymentRepository) *PaymentService {
	return &PaymentService{repo}
}

func (s *PaymentService) CreatePayment(
	ctx context.Context,
	orderID string,
	amount float64,
) (*models.PaymentModel, error) {

	if amount <= 0 {
		return nil, errors.New("amount must be greater than zero")
	}

	payment := &models.PaymentModel{
		ID:          uuid.New().String(),
		OrderID:     orderID,
		Amount:      amount,
		Status:      "success", // simulate gateway success
		PaymentDate: time.Now(),
	}

	if err := s.repo.Create(ctx, payment); err != nil {
		return nil, err // duplicate payment → DB blocks
	}

	return payment, nil
}
