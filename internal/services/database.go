package services

import (
	"os"
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
		// ✅ FIXED: Use PostgreSQL connection string from environment
		dbAddr := os.Getenv("DB_ADDR")
		if dbAddr == "" {
			// Default to PostgreSQL with correct password rafa2005
			dbAddr = "postgres://postgres:rafa2005@localhost:5432/reloop?sslmode=disable"
		}

		database, err := db.New(
			dbAddr, // PostgreSQL connection string, not SQLite file path
			30, 30, "15m",
		)
		if err != nil {
			panic("Failed to connect to database: " + err.Error())
		}

		dbInstance = &DatabaseManager{db: database}
	})
	return dbInstance
}

func (dm *DatabaseManager) GetDB() *gorm.DB {
	return dm.db
}
