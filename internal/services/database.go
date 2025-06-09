package services

import (
	"reloop-backend/internal/db"
	"sync"

	"gorm.io/gorm"
)

type DatabaseManager struct {
	db *gorm.DB
}

var (
	dbInstance *DatabaseManager
	once       sync.Once
)

func GetDatabase() *DatabaseManager {
	once.Do(func() {
		// Menggunakan config yang sama dari main.go
		database, err := db.New(
			"postgres://postgres:1234@localhost:5432/reloop?sslmode=disable",
			30, 30, "15m",
		)
		if err != nil {
			panic("Failed to connect to database")
		}
		dbInstance = &DatabaseManager{db: database}
	})
	return dbInstance
}

func (dm *DatabaseManager) GetDB() *gorm.DB {
	return dm.db
}
