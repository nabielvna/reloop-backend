package store

import (
	"context"

	"gorm.io/gorm"
)

type Storage struct {
	Users interface {
		Create(context.Context, *User) error
	}
}

func NewStorage(db *gorm.DB) Storage {
	return Storage{
		Users: &UsersStore{db},
	}
}