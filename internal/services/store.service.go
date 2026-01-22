package services

import (
	"context"
	"errors"
	"ecommerce/internal/models"
	"ecommerce/internal/repositories"

	"github.com/google/uuid"
)

type StoreService struct {
	repo repositories.StoreRepository
}

func NewStoreService(repo repositories.StoreRepository) *StoreService {
	return &StoreService{repo}
}

func (s *StoreService) CreateStore(ctx context.Context, store *models.StoreModel) error {
	if store.Name == "" {
		return errors.New("store name is required")
	}

	if store.Status == "" {
		store.Status = "active"
	}

	store.ID = uuid.New().String()
	return s.repo.Create(ctx, store)
}

func (s *StoreService) GetStores(ctx context.Context) ([]models.StoreModel, error) {
	return s.repo.GetAll(ctx)
}
