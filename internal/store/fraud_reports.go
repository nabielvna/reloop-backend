package store

import "gorm.io/gorm"

type FraudReport struct {
	gorm.Model
	ReporterID     *uint
	ReportedItemID *uint
	ReportDetails  string `gorm:"not null"`
	Status         string `gorm:"type:varchar(20);not null;default:submitted"`

	Reporter     *User `gorm:"foreignKey:ReporterID"`
	ReportedItem *Item `gorm:"foreignKey:ReportedItemID"`
}