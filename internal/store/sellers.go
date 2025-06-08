package store

import "gorm.io/gorm"

type SellersStore struct {
	db *gorm.DB
}

func NewSellersStore(db *gorm.DB) *SellersStore {
    return &SellersStore{
        db: db,
    }
}