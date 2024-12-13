package main

import (
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	r := chi.NewRouter()

	// Global middlewares
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Inject database pool into context
	r.Use(app.DBMiddleware(app.DB))

	// Serve static files from the Astro build directory
	buildPath := "./frontend/dist"
	fs := http.FileServer(http.Dir(buildPath))
	r.Handle("/static/*", fs)

	// SPA fallback to serve index.html for unknown routes
	r.NotFound(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join(buildPath, "index.html"))
	}))

	// Mount authenticated API router
	r.Mount("/api", app.AuthenticatedAPIRouter())

	return r
}

// AuthenticatedAPIRouter creates a router for authenticated API routes
func (app *application) AuthenticatedAPIRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(app.AuthMiddleware) // Apply authentication middleware to all routes under /api

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the Authenticated API"))
	})

	// Example: Add additional authenticated API routes
	r.Get("/profile", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Your profile details"))
	})

	r.Get("/settings", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Your account settings"))
	})

	return r
}
