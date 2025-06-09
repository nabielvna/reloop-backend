package interfaces

import (
	"context"
	"reloop-backend/internal/models"
)

type SellerRepositoryInterface interface {
	Create(ctx context.Context, seller *models.Seller) error
	GetByID(ctx context.Context, id uint) (*models.Seller, error)
	GetByUserID(ctx context.Context, userID uint) (*models.Seller, error)
	Update(ctx context.Context, seller *models.Seller) error
	Delete(ctx context.Context, id uint) error
	UpdateVerificationStatus(ctx context.Context, id uint, status string) error
	GetPendingVerification(ctx context.Context) ([]models.Seller, error)
	GetAll(ctx context.Context) ([]models.Seller, error)
}
