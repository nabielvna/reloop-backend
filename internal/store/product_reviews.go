package store

import "gorm.io/gorm"

type ProductReviews struct {
	db *gorm.DB
}

func NewProductReviewsStore(db *gorm.DB) *ProductReviews {
	return &ProductReviews{
		db: db,
	}
}