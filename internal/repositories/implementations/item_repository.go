package implementations

import (
	"context"
	"reloop-backend/internal/models"
	"reloop-backend/internal/repositories/interfaces"
	"reloop-backend/internal/services"

	"gorm.io/gorm"
)

type ItemRepository struct {
	db *gorm.DB
}

func NewItemRepository() interfaces.ItemRepositoryInterface {
	db := services.GetDatabase().GetDB()
	return &ItemRepository{db: db}
}

func (r *ItemRepository) Create(ctx context.Context, item *models.Item) error {
	return r.db.WithContext(ctx).Create(item).Error
}

func (r *ItemRepository) GetByID(ctx context.Context, id uint) (*models.Item, error) {
	var item models.Item
	err := r.db.WithContext(ctx).
		Preload("Seller").
		Preload("Seller.User").
		Preload("Category").
		First(&item, id).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *ItemRepository) GetBySeller(ctx context.Context, sellerID uint) ([]models.Item, error) {
	var items []models.Item
	err := r.db.WithContext(ctx).
		Preload("Category").
		Where("seller_id = ?", sellerID).
		Find(&items).Error
	return items, err
}

func (r *ItemRepository) Browse(ctx context.Context, filters interfaces.ItemFilters) ([]models.Item, error) {
	var items []models.Item
	query := r.db.WithContext(ctx).
		Preload("Seller").
		Preload("Seller.User").
		Preload("Category")

	// Apply filters
	if filters.CategoryID != nil {
		query = query.Where("category_id = ?", *filters.CategoryID)
	}
	if filters.MinPrice != nil {
		query = query.Where("price >= ?", *filters.MinPrice)
	}
	if filters.MaxPrice != nil {
		query = query.Where("price <= ?", *filters.MaxPrice)
	}
	if filters.Status != nil {
		query = query.Where("status = ?", *filters.Status)
	}
	if filters.Search != nil {
		query = query.Where("name ILIKE ? OR description ILIKE ?",
			"%"+*filters.Search+"%", "%"+*filters.Search+"%")
	}

	err := query.Find(&items).Error
	return items, err
}

func (r *ItemRepository) Update(ctx context.Context, item *models.Item) error {
	return r.db.WithContext(ctx).Save(item).Error
}

func (r *ItemRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Item{}, id).Error
}

func (r *ItemRepository) UpdateStatus(ctx context.Context, id uint, status string) error {
	return r.db.WithContext(ctx).Model(&models.Item{}).Where("id = ?", id).Update("status", status).Error
}

func (r *ItemRepository) GetAll(ctx context.Context) ([]models.Item, error) {
	var items []models.Item
	err := r.db.WithContext(ctx).
		Preload("Seller").
		Preload("Seller.User").
		Preload("Category").
		Find(&items).Error
	return items, err
}
