package store

import (
	"context"
	"reloop-backend/internal/models"

	"gorm.io/gorm"
)

type UsersStore struct {
    db *gorm.DB
}

func NewUsersStore(db *gorm.DB) *UsersStore {
    return &UsersStore{
        db: db,
    }
}

func (s *UsersStore) Create(ctx context.Context, user *models.User) error {
    return s.db.WithContext(ctx).Create(user).Error
}

func (s *UsersStore) GetByID(ctx context.Context, id uint) (*models.User, error) {
    var user models.User
    if err := s.db.WithContext(ctx).First(&user, id).Error; err != nil {
        return nil, err
    }
    return &user, nil
}

func (s *UsersStore) GetByEmail(ctx context.Context, email string) (*models.User, error) {
    var user models.User
    if err := s.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
        return nil, err
    }
    return &user, nil
}

func (s *UsersStore) Update(ctx context.Context, user *models.User) error {
    return s.db.WithContext(ctx).Save(user).Error
}

func (s *UsersStore) Delete(ctx context.Context, id uint) error {
    return s.db.WithContext(ctx).Delete(&models.User{}, id).Error
}