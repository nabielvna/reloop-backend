package db

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func New(addr string, maxOpenConns, maxIdleConns int, maxIdleTime string) (*gorm.DB, error) {
	// For SQLite, addr is the file path
	db, err := gorm.Open(sqlite.Open(addr), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(maxOpenConns)
	sqlDB.SetMaxIdleConns(maxIdleConns)

	duration, err := time.ParseDuration(maxIdleTime)
	if err != nil {
		return nil, err
	}

	sqlDB.SetConnMaxIdleTime(duration)

	return db, nil
}
