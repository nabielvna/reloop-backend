package store

import "gorm.io/gorm"

type AdminsStore struct {
	db *gorm.DB
}

func NewAdminsStore(db *gorm.DB) *AdminsStore {
	return &AdminsStore{
		db: db,
	}
}