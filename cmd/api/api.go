package main

import (
	"net/http"
	"reloop-backend/internal/store"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	_ "reloop-backend/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

type application struct {
	config config
	store  *store.Storage
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

	r.Get("/swagger/*", httpSwagger.WrapHandler)

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)

		// Public routes
		r.Post("/auth/register", app.registerUserHandler)
		r.Post("/auth/login", app.loginUserHandler)

		// Public item browsing
		r.Get("/items", app.browseItemsHandler)
		r.Get("/items/{itemID}", app.getItemHandler)

		// Protected routes
		r.Group(func(r chi.Router) {
			r.Use(app.authMiddleware)

			// User profile
			r.Get("/me", app.getUserProfileHandler)

			// Item management (for sellers)
			r.Post("/items", app.createItemHandler)
			r.Get("/my-items", app.getMyItemsHandler)
			r.Put("/items/{itemID}", app.updateItemHandler)
			r.Delete("/items/{itemID}", app.deleteItemHandler)

			// Admin only routes
			r.Route("/admin", func(r chi.Router) {
				// TODO: Add admin middleware
				r.Patch("/items/{itemID}/status", app.updateItemStatusHandler)
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

	return srv.ListenAndServe()
}
