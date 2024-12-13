package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/iankencruz/eggcounter/backend/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

// LoggerMiddleware logs all HTTP requests
func (app *application) LoggerMiddleware(next http.Handler) http.Handler {
	return middleware.Logger(next)
}

// RecoveryMiddleware recovers from panics
func (app *application) RecoveryMiddleware(next http.Handler) http.Handler {
	return middleware.Recoverer(next)
}

// DBMiddleware injects the database connection pool into the context
func (app *application) DBMiddleware(db *pgxpool.Pool) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r = r.WithContext(models.WithDB(r.Context(), db))
			next.ServeHTTP(w, r)
		})
	}
}

// AuthMiddleware checks if the user is authenticated and redirects to /login if not
func (app *application) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, err := app.User.GetSessionUserID(r)
		if err != nil || userID == 0 {
			log.Println("AuthMiddleware error: ", err) // Log the error for debugging
			http.Redirect(w, r, "/auth/login", http.StatusFound)
			return
		}

		next.ServeHTTP(w, r)
	})
}
