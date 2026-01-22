package services

import (
	"context"
	"ecommerce/internal/models"
	"ecommerce/internal/repositories"

	"github.com/google/uuid"
)

type UserService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *UserService {
	return &UserService{repo}
}

func (s *UserService) CreateUser(ctx context.Context, user *models.UserModel) error {
	user.ID = uuid.New().String()
	return s.repo.Create(ctx, user)
}

func (s *UserService) GetUser(ctx context.Context, id string) (*models.UserModel, error) {
	return s.repo.GetByID(ctx, id)
}
