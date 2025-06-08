package store

import "gorm.io/gorm"

type CategoriesStore struct {
	db *gorm.DB
}

func NewCategoriesStore(db *gorm.DB) *CategoriesStore {
	return &CategoriesStore{
		db: db,
	}
}