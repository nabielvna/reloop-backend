package implementations

import (
	"context"
	"reloop-backend/internal/models"
	"reloop-backend/internal/repositories/interfaces"

	// "reloop-backend/internal/services" // Hapus ini

	"gorm.io/gorm"
)

type SellerRepository struct {
	db *gorm.DB
}

func NewSellerRepository(db *gorm.DB) interfaces.SellerRepositoryInterface {
	return &SellerRepository{db: db}
}

func (r *SellerRepository) Create(ctx context.Context, seller *models.Seller) error {
	return r.db.WithContext(ctx).Create(seller).Error
}

func (r *SellerRepository) GetByID(ctx context.Context, id uint) (*models.Seller, error) {
	var seller models.Seller
	err := r.db.WithContext(ctx).
		Preload("User").
		Preload("Items").
		First(&seller, id).Error
	if err != nil {
		return nil, err
	}
	return &seller, nil
}

func (r *SellerRepository) GetByUserID(ctx context.Context, userID uint) (*models.Seller, error) {
	var seller models.Seller
	err := r.db.WithContext(ctx).
		Preload("User").
		Preload("Items").
		Where("user_id = ?", userID).
		First(&seller).Error
	if err != nil {
		return nil, err
	}
	return &seller, nil
}

func (r *SellerRepository) Update(ctx context.Context, seller *models.Seller) error {
	return r.db.WithContext(ctx).Save(seller).Error
}

func (r *SellerRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Seller{}, id).Error
}

func (r *SellerRepository) UpdateVerificationStatus(ctx context.Context, id uint, status string) error {
	return r.db.WithContext(ctx).
		Model(&models.Seller{}).
		Where("user_id = ?", id).
		Update("verification_status", status).Error
}

func (r *SellerRepository) GetPendingVerification(ctx context.Context) ([]models.Seller, error) {
	var sellers []models.Seller
	err := r.db.WithContext(ctx).
		Preload("User").
		Where("verification_status = ?", "pending").
		Find(&sellers).Error
	return sellers, err
}

func (r *SellerRepository) GetAll(ctx context.Context) ([]models.Seller, error) {
	var sellers []models.Seller
	err := r.db.WithContext(ctx).
		Preload("User").
		Find(&sellers).Error
	return sellers, err
}