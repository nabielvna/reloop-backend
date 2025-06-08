package store

import "github.com/lib/pq"

type Admin struct {
	UserID      uint           `gorm:"primaryKey"`
	Permissions pq.StringArray `gorm:"type:text[]"`

	User *User `gorm:"foreignKey:UserID"`
}