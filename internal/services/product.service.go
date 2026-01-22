package services

import (
	"context"
	"errors"
	"ecommerce/internal/models"
	"ecommerce/internal/repositories"

	"github.com/google/uuid"
)

type ProductService struct {
	repo repositories.ProductRepository
}

func NewProductService(repo repositories.ProductRepository) *ProductService {
	return &ProductService{repo}
}

func (s *ProductService) CreateProduct(
	ctx context.Context,
	storeID string,
	product *models.ProductModel,
) error {

	if product.Name == "" {
		return errors.New("product name is required")
	}

	if product.Price < 0 {
		return errors.New("price cannot be negative")
	}

	product.ID = uuid.New().String()
	product.StoreID = storeID

	return s.repo.Create(ctx, product)
}

func (s *ProductService) GetProductByID(
	ctx context.Context,
	id string,
) (*models.ProductModel, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *ProductService) GetProductsByStore(
	ctx context.Context,
	storeID string,
) ([]models.ProductModel, error) {
	return s.repo.GetByStore(ctx, storeID)
}
