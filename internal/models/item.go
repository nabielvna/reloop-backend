package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	SellerID    uint   `gorm:"not null"`
	CategoryID  uint   `gorm:"not null"`
	Name        string `gorm:"not null;size:255"`
	Description string
	Price       float64 `gorm:"type:numeric(12,2);not null"`
	Quantity    int     `gorm:"not null"`
	Status      string  `gorm:"type:varchar(10);not null;default:pending"`

	Seller   *Seller   `gorm:"foreignKey:SellerID"`
	Category *Category `gorm:"foreignKey:CategoryID"`

	ProductReviews []ProductReview `gorm:"foreignKey:ItemID"`
	FraudReports   []FraudReport   `gorm:"foreignKey:ReportedItemID"`
}