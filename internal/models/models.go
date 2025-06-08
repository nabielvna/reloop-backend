package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model 
	Email        string    `gorm:"unique;not null;size:255"`
	PasswordHash string    `gorm:"not null;size:255"`
	UserName     string    `gorm:"not null;size:100"`
	Role         string    `gorm:"type:varchar(10);not null;default:user"`

	Seller *Seller `gorm:"foreignKey:UserID"`
	Admin  *Admin  `gorm:"foreignKey:UserID"`
	
	ProductReviews []ProductReview `gorm:"foreignKey:UserID"`
	FraudReports   []FraudReport   `gorm:"foreignKey:ReporterID"`
}

type Seller struct {
	UserID             uint   `gorm:"primaryKey"`
	BusinessName       string `gorm:"not null;size:255"`
	WhatsappNumber     string `gorm:"not null;size:25"`
	WhatsappLink       string `gorm:"size:255"`
	ProfilePictureURL  string `gorm:"size:255"`
	VerificationStatus string `gorm:"size:50"`
	AccountStatus      string `gorm:"type:varchar(25);not null;default:pending_verification"`

	User  *User  `gorm:"foreignKey:UserID"`
	Items []Item `gorm:"foreignKey:SellerID"`
}

type Admin struct {
	UserID      uint           `gorm:"primaryKey"`
	Permissions pq.StringArray `gorm:"type:text[]"`

	User *User `gorm:"foreignKey:UserID"`
}

type Category struct {
	gorm.Model 
	Name             string    `gorm:"not null;size:100"`
	ParentCategoryID *uint     
	IsActive         bool      `gorm:"not null;default:true"`

	Items []Item `gorm:"foreignKey:CategoryID"`
	
	SubCategories  []Category `gorm:"foreignKey:ParentCategoryID"`
	ParentCategory *Category  `gorm:"foreignKey:ParentCategoryID"`
}

type Item struct {
	gorm.Model 
	SellerID    uint      `gorm:"not null"`
	CategoryID  uint      `gorm:"not null"`
	Name        string    `gorm:"not null;size:255"`
	Description string
	Price       float64   `gorm:"type:numeric(12,2);not null"`
	Quantity    int       `gorm:"not null"`
	Status      string    `gorm:"type:varchar(10);not null;default:pending"`

	Seller   *Seller   `gorm:"foreignKey:SellerID"`
	Category *Category `gorm:"foreignKey:CategoryID"`

	ProductReviews []ProductReview `gorm:"foreignKey:ItemID"`
	FraudReports   []FraudReport   `gorm:"foreignKey:ReportedItemID"`
}

type FraudReport struct {
	gorm.Model 
	ReporterID     *uint     
	ReportedItemID *uint     
	ReportDetails  string    `gorm:"not null"`
	Status         string    `gorm:"type:varchar(20);not null;default:submitted"`
	
	Reporter     *User `gorm:"foreignKey:ReporterID"`
	ReportedItem *Item `gorm:"foreignKey:ReportedItemID"`
}

type ProductReview struct {
	gorm.Model 
	ItemID    uint      `gorm:"uniqueIndex:idx_item_user_review;not null"`
	UserID    uint      `gorm:"uniqueIndex:idx_item_user_review;not null"`
	Rating    int       `gorm:"not null"`
	Comment   string

	Item *Item `gorm:"foreignKey:ItemID"`
	User *User `gorm:"foreignKey:UserID"`
}