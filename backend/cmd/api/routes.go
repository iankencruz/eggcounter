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

	// Load and save session middleware for all routes
	router.Use(app.Session.LoadAndSave)

	// 🔐 Public routes
	router.Post("/api/register", app.registerHandler)
	router.Post("/api/login", app.loginHandler)
	router.Post("/api/logout", app.logoutHandler)
	router.Get("/api/auth/status", app.authStatusHandler)

	// 🔒 Protected API routes (Require Auth)
	router.Route("/api", func(r chi.Router) {
		r.Use(app.requireAuth) // Ensure all these routes require authentication

		// 🥚 Egg Routes
		r.Get("/dashboard", app.dashboardHandler)
		r.Get("/eggcount", app.getEggCountHandler)         // Fetch total egg count
		r.Post("/eggcount", app.addEggCountHandler)        // Add egg count
		r.Delete("/eggcount/{id}", app.deleteEntryHandler) // Delete an egg count entry

		// 👫 Friends Routes (Nested Group)
		r.Route("/friends", func(fr chi.Router) {
			fr.Post("/requests", app.sendFriendRequestHandler)      // Send a friend request
			fr.Get("/requests", app.getFriendRequestsHandler)       // View pending friend requests
			fr.Post("/accept/{id}", app.acceptFriendRequestHandler) // Accept a friend request by id
			fr.Post("/reject/{id}", app.rejectFriendRequestHandler) // Reject a friend request by id
			fr.Get("/", app.getFriendsListHandler)                  // Get all friends for the user
		})
	})

	// 🗂️ Get the project root directory
	rootPath, err := os.Getwd()
	if err != nil {
		panic("Unable to get current working directory: " + err.Error())
	}

	// 🗂️ Static files path
	staticPath := filepath.Join(rootPath, "frontend", "build")
	log.Printf("Serving static files from %s", staticPath)

	// 📁 Serve static files
	fileServer := http.FileServer(http.Dir(staticPath))
	router.Handle("/static/*", http.StripPrefix("/static", fileServer))
	router.Handle("/favicon.ico", http.StripPrefix("/", fileServer))
	router.Handle("/manifest.json", http.StripPrefix("/", fileServer))

	// 🔍 Custom 404 Page for Unmatched Routes
	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		// Avoid matching API routes with the 404 fallback
		if isAPIRequest(r.URL.Path) {
			http.NotFound(w, r)
			return
		}
		// Serve 404.html file
		log.Printf("404 Not Found: %s", r.URL.Path)
		w.WriteHeader(http.StatusNotFound)
		http.ServeFile(w, r, filepath.Join(staticPath, "404.html"))
	})

	// 🕵️‍♂️ Fallback handler for SPA routes
	router.Handle("/*", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestedPath := filepath.Join(staticPath, r.URL.Path)
		if fileExists(requestedPath) {
			http.ServeFile(w, r, requestedPath)
		} else if !isAPIRequest(r.URL.Path) { // Avoid falling back for API routes
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

// Helper function to detect if the route is an API request
func isAPIRequest(path string) bool {
	return len(path) >= 4 && path[:4] == "/api"
}
