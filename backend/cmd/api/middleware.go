package main

import (
	"encoding/json"
	"net/http"
)

func (app *Application) requireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := app.Session.GetInt(r.Context(), "userID")
		if userID == 0 {
			// Send a 401 response without redirecting
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Unauthorized. Please log in.",
			})
			return
		}
		next.ServeHTTP(w, r)
	})
}
