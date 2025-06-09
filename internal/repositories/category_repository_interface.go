package interfaces

import (
	"context"
	"reloop-backend/internal/models"
)

type CategoryRepositoryInterface interface {
	Create(ctx context.Context, category *models.Category) error
	GetByID(ctx context.Context, id uint) (*models.Category, error)
	GetAll(ctx context.Context) ([]models.Category, error)
	GetByParent(ctx context.Context, parentID uint) ([]models.Category, error)
	Update(ctx context.Context, category *models.Category) error
	Delete(ctx context.Context, id uint) error
	GetActive(ctx context.Context) ([]models.Category, error)
}
