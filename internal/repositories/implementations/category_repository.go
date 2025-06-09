package implementations

import (
	"context"
	"reloop-backend/internal/models"
	"reloop-backend/internal/repositories/interfaces"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) interfaces.CategoryRepositoryInterface {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) Create(ctx context.Context, category *models.Category) error {
	return r.db.WithContext(ctx).Create(category).Error
}

func (r *CategoryRepository) GetByID(ctx context.Context, id uint) (*models.Category, error) {
	var category models.Category
	err := r.db.WithContext(ctx).
		Preload("SubCategories").
		Preload("ParentCategory").
		First(&category, id).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *CategoryRepository) GetAll(ctx context.Context) ([]models.Category, error) {
	var categories []models.Category
	err := r.db.WithContext(ctx).
		Preload("SubCategories").
		Preload("ParentCategory").
		Find(&categories).Error
	return categories, err
}

func (r *CategoryRepository) GetByParent(ctx context.Context, parentID uint) ([]models.Category, error) {
	var categories []models.Category
	err := r.db.WithContext(ctx).
		Where("parent_category_id = ?", parentID).
		Find(&categories).Error
	return categories, err
}

func (r *CategoryRepository) Update(ctx context.Context, category *models.Category) error {
	return r.db.WithContext(ctx).Save(category).Error
}

func (r *CategoryRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Category{}, id).Error
}

func (r *CategoryRepository) GetActive(ctx context.Context) ([]models.Category, error) {
	var categories []models.Category
	err := r.db.WithContext(ctx).
		Where("is_active = ?", true).
		Find(&categories).Error
	return categories, err
}