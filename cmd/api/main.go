package main

import (
	"log"
	"reloop-backend/internal/db"
	"reloop-backend/internal/env"
	"reloop-backend/internal/models"
	"reloop-backend/internal/store"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  .env file tidak ditemukan, lanjut dengan env vars sistem atau fallback")
	}
	
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr: env.GetString("DB_ADDR", "postgres://postgres:1234@localhost:5432/reloop?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
	}

	db, err := db.New(
		cfg.db.addr, 
		cfg.db.maxOpenConns, 
		cfg.db.maxIdleConns, 
		cfg.db.maxIdleTime,
	)

	db.AutoMigrate(
		&models.User{},
	)

	if err != nil {
		log.Panic(err)
	}

	log.Printf("→ DB_ADDR = %s\n", cfg.db.addr)

	log.Println("Koneksi ke database berhasil")
	
	store := store.NewStorage(db)

	app := &application{
		config: cfg,
		store: store,
	}


	mux := app.mount()

	log.Fatal(app.run(mux))	
}