package store

import "gorm.io/gorm"

type ProductReview struct {
	gorm.Model
	ItemID  uint `gorm:"uniqueIndex:idx_item_user_review;not null"`
	UserID  uint `gorm:"uniqueIndex:idx_item_user_review;not null"`
	Rating  int  `gorm:"not null"`
	Comment string

	Item *Item `gorm:"foreignKey:ItemID"`
	User *User `gorm:"foreignKey:UserID"`
}