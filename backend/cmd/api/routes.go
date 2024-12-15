package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

func (app *Application) routes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(app.Session.LoadAndSave)
	// Public routes
	router.Post("/api/register", app.registerHandler)
	router.Post("/api/login", app.loginHandler)
	router.Post("/api/logout", app.logoutHandler)
	router.Get("/api/auth/status", app.authStatusHandler)

	// Protected routes
	router.Route("/api", func(r chi.Router) {
		r.Use(app.Session.LoadAndSave)
		r.Use(app.requireAuth)
		r.Get("/dashboard", app.dashboardHandler)
	})

	// Get the project root directory
	rootPath, err := os.Getwd()
	if err != nil {
		panic("Unable to get current working directory: " + err.Error())
	}

	// Static files path
	staticPath := filepath.Join(rootPath, "frontend", "build")

	// Log where files are served from
	log.Printf("Serving static files from %s", staticPath)

	// Serve static files
	fileServer := http.FileServer(http.Dir(staticPath))
	router.Handle("/static/*", http.StripPrefix("/static", fileServer))
	router.Handle("/favicon.ico", http.StripPrefix("/", fileServer))
	router.Handle("/manifest.json", http.StripPrefix("/", fileServer))

	// Handle unmatched routes with a custom 404 page
	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		// Serve 404.html file
		w.WriteHeader(http.StatusNotFound)
		http.ServeFile(w, r, filepath.Join(staticPath, "404.html"))
	})

	// Fallback handler
	router.Handle("/*", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestedPath := filepath.Join(staticPath, r.URL.Path)
		if fileExists(requestedPath) {
			http.ServeFile(w, r, requestedPath)
		} else if r.URL.Path != "/api/dashboard" { // Avoid falling back for API routes
			log.Printf("Falling back to index.html for route: %s", r.URL.Path)
			http.ServeFile(w, r, filepath.Join(staticPath, "index.html"))
		} else {
			http.NotFound(w, r)
		}
	}))

	return router
}

// Helper function to check if a file exists
func fileExists(filePath string) bool {
	info, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
