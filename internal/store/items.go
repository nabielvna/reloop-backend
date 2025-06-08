package store

import "gorm.io/gorm"

type ItemsStore struct {
	db *gorm.DB
}

func NewItemsStore(db *gorm.DB) *ItemsStore {
	return &ItemsStore{
		db: db,
	}
}