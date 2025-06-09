package main

import (
	"log"
	"reloop-backend/internal/env"
	"reloop-backend/internal/models"
	"reloop-backend/internal/services"
	"reloop-backend/internal/store"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(" .env file tidak ditemukan, lanjut dengan env vars sistem atau fallback")
	}

	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			// pass : rafael2005
			addr:         env.GetString("DB_ADDR", "postgres://postgres:rafael2005@localhost:5432/reloop?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		jwtSecret: env.GetString("JWT_SECRET", "secret"),
	}

	// Initialize database singleton
	db := services.GetDatabase().GetDB()

	// Auto migrate with better error handling
	err := db.AutoMigrate(
		&models.Admin{},
		&models.Category{},
		&models.FraudReport{},
		&models.Item{},
		&models.ProductReview{},
		&models.Seller{},
		&models.User{},
	)
	if err != nil {
		log.Fatalf(" Migration failed: %v", err)
	}

	log.Printf("→ DB_ADDR = %s\n", cfg.db.addr)
	log.Println("Koneksi ke database berhasil")
	store := store.NewStorage(cfg.jwtSecret)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()
	log.Printf("🚀 Server starting on %s", cfg.addr)
	log.Fatal(app.run(mux))
}
