package services

import (
	"context"
	"errors"
	"time"

	"ecommerce/internal/models"
	"ecommerce/internal/repositories"

	"github.com/google/uuid"
)

type SubscriptionService struct {
	repo repositories.SubscriptionRepository
}

func NewSubscriptionService(repo repositories.SubscriptionRepository) *SubscriptionService {
	return &SubscriptionService{repo}
}

func (s *SubscriptionService) CreateSubscription(
	ctx context.Context,
	sub *models.SubscriptionModel,
) error {

	if sub.UserID == "" || sub.ProductID == "" {
		return errors.New("user_id and product_id are required")
	}

	sub.ID = uuid.New().String()
	sub.StartDate = time.Now()

	if sub.Status == "" {
		sub.Status = "active"
	}

	return s.repo.Create(ctx, sub)
}

func (s *SubscriptionService) GetUserSubscriptions(
	ctx context.Context,
	userID string,
) ([]models.SubscriptionModel, error) {
	return s.repo.GetByUser(ctx, userID)
}
