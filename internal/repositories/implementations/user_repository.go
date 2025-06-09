package implementations

import (
	"context"
	"reloop-backend/internal/models"
	"reloop-backend/internal/repositories/interfaces"
	"reloop-backend/internal/services"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() interfaces.UserRepositoryInterface {
	db := services.GetDatabase().GetDB()
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user *models.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *UserRepository) GetByID(ctx context.Context, id uint) (*models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).Preload("Seller").Preload("Admin").First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Update(ctx context.Context, user *models.User) error {
	return r.db.WithContext(ctx).Save(user).Error
}

func (r *UserRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.User{}, id).Error
}

func (r *UserRepository) GetAll(ctx context.Context) ([]models.User, error) {
	var users []models.User
	err := r.db.WithContext(ctx).Preload("Seller").Preload("Admin").Find(&users).Error
	return users, err
}

func (r *UserRepository) UpdateRole(ctx context.Context, id uint, role string) error {
	return r.db.WithContext(ctx).Model(&models.User{}).Where("id = ?", id).Update("role", role).Error
}

func (r *UserRepository) UpdateStatus(ctx context.Context, id uint, status string) error {
	// Note: User model doesn't have status field, but we'll add it if needed
	return r.db.WithContext(ctx).Model(&models.User{}).Where("id = ?", id).Update("status", status).Error
}
