package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email        string `gorm:"unique;not null;size:255"`
	PasswordHash string `gorm:"not null;size:255"`
	UserName     string `gorm:"not null;size:100"`
	Role         string `gorm:"type:varchar(10);not null;default:user"`

	Seller *Seller `gorm:"foreignKey:UserID"`
	Admin  *Admin  `gorm:"foreignKey:UserID"`

	ProductReviews []ProductReview `gorm:"foreignKey:UserID"`
	FraudReports   []FraudReport   `gorm:"foreignKey:ReporterID"`
}