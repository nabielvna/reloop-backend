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

func GetDatabase(dsn string, maxOpenConns, maxIdleConns int, maxIdleTime string) *DatabaseManager {
	once.Do(func() {
		database, err := db.New(
			dsn,
			maxOpenConns,
			maxIdleConns,
			maxIdleTime,
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