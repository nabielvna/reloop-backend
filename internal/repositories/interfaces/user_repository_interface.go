package interfaces

import (
	"context"
	"reloop-backend/internal/models"
)

type UserRepositoryInterface interface {
	Create(ctx context.Context, user *models.User) error
	GetByID(ctx context.Context, id uint) (*models.User, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id uint) error
	GetAll(ctx context.Context) ([]models.User, error)
	UpdateRole(ctx context.Context, id uint, role string) error
	UpdateStatus(ctx context.Context, id uint, status string) error
}
