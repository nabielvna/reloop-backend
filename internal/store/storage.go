package store

import (
	"context"
	"reloop-backend/internal/models"

	"gorm.io/gorm"
)

type Storage struct {
	Users interface {
		Create(context.Context, *models.User) error
	}
}

func NewStorage(db *gorm.DB) Storage {
	return Storage{
		Users: &UsersStore{db},
	}
}