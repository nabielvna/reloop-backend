package store

import "gorm.io/gorm"

type FraudReportsStore struct {
	db *gorm.DB
}

func NewFraudReportsStore(db *gorm.DB) *FraudReportsStore {
	return &FraudReportsStore{
		db: db,
	}
}