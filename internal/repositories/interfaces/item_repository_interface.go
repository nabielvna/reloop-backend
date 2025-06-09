package interfaces

import (
	"context"
	"reloop-backend/internal/models"
)

type ItemFilters struct {
	CategoryID *uint
	MinPrice   *float64
	MaxPrice   *float64
	Condition  *string
	Location   *string
	Search     *string
	Status     *string
	SellerID   *uint
}

type ItemRepositoryInterface interface {
	Create(ctx context.Context, item *models.Item) error
	GetByID(ctx context.Context, id uint) (*models.Item, error)
	GetBySeller(ctx context.Context, sellerID uint) ([]models.Item, error)
	Browse(ctx context.Context, filters ItemFilters) ([]models.Item, error)
	Update(ctx context.Context, item *models.Item) error
	Delete(ctx context.Context, id uint) error
	UpdateStatus(ctx context.Context, id uint, status string) error
	GetAll(ctx context.Context) ([]models.Item, error)
}
