package repositories

import (
	"context"
	"database/sql"
	"ecommerce/internal/models"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.UserModel) error
	GetByID(ctx context.Context, id string) (*models.UserModel, error)
}

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) UserRepository {
	return &userRepo{db}
}

func (r *userRepo) Create(ctx context.Context, user *models.UserModel) error {
	query := `INSERT INTO users (id, name, email, phone)
	          VALUES ($1, $2, $3, $4)`
	_, err := r.db.ExecContext(ctx, query,
		user.ID, user.Name, user.Email, user.Phone)
	return err
}

func (r *userRepo) GetByID(ctx context.Context, id string) (*models.UserModel, error) {
	query := `SELECT id, name, email, phone FROM users WHERE id=$1`
	row := r.db.QueryRowContext(ctx, query, id)

	var user models.UserModel
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Phone)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
