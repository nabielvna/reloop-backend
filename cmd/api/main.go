// @title Reloop API
// @version 1.0
// @description Ini adalah dokumentasi API untuk backend Reloop.

// @host localhost:8080
// @BasePath /v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Ketik "Bearer" diikuti spasi dan token JWT Anda.

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
			addr:         env.GetString("DB_ADDR", "postgres://postgres:1234@localhost:5432/reloop?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		jwtSecret: env.GetString("JWT_SECRET", "secret"),
	}

	db := services.GetDatabase(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	).GetDB()

	db.AutoMigrate(
		&models.Admin{},
		&models.Category{},
		&models.FraudReport{},
		&models.Item{},
		&models.ProductReview{},
		&models.Seller{},
		&models.User{},
	)

	log.Printf("→ DB_ADDR = %s\n", cfg.db.addr)
	log.Println("Koneksi ke database berhasil")

	store := store.NewStorage(db, cfg.jwtSecret)
	app := &application{
		config: cfg,
		store:  store,
	}

	log.Printf("Menjalankan server di %s", cfg.addr)
	mux := app.mount()
	log.Printf("🚀 Server starting on %s", cfg.addr)
	log.Fatal(app.run(mux))
}