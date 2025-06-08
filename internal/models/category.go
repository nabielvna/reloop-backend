package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name             string `gorm:"not null;size:100"`
	ParentCategoryID *uint
	IsActive         bool `gorm:"not null;default:true"`

	Items []Item `gorm:"foreignKey:CategoryID"`

	SubCategories  []Category `gorm:"foreignKey:ParentCategoryID"`
	ParentCategory *Category  `gorm:"foreignKey:ParentCategoryID"`
}