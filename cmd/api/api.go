package main

import (
	"encoding/json"
	"log"
	"net/http"
	"reloop-backend/internal/store"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	config config
	store  store.Storage
}

type config struct {
	addr      string
	db        dbConfig
	jwtSecret string
}

type dbConfig struct {
	addr         string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

func (app *application) mount() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)

		r.Post("/users/register", app.registerUserHandler)
		r.Post("/users/login", app.loginUserHandler)

		// Rute untuk otentikasi
		r.Group(func(r chi.Router) {
			r.Use(app.authMiddleware)
			r.Get("/me", func(w http.ResponseWriter, r *http.Request) {
				userID, ok := r.Context().Value(userContextKey).(uint)
				if !ok {
					http.Error(w, "Gagal mendapatkan user dari context", http.StatusInternalServerError)
					return
				}
				user, err := app.store.Users.GetByID(r.Context(), userID)
				if err != nil {
					http.Error(w, "User tidak ditemukan", http.StatusNotFound)
					return
				}

				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(user)
			})
		})
	})

	return r
}

func (app *application) run(mux *chi.Mux) error {

	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("Starting server on %s", app.config.addr)

	return srv.ListenAndServe()
}
