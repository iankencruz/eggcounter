package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"github.com/iankencruz/eggcounter/backend/internal/models"
)

func (app *Application) dashboardHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve user ID from session
	userID := app.Session.GetInt(r.Context(), "userID")
	if userID == 0 {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Unauthorized. Please log in.",
			"status":  http.StatusUnauthorized,
		})
		return
	}

	// Fetch user details
	user, err := app.UserModel.GetUserByID(r.Context(), userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Failed to retrieve user data",
			"status":  http.StatusInternalServerError,
		})
		return
	}

	// Build the dashboard response
	data := map[string]interface{}{
		"user": map[string]string{
			"firstName": user.FirstName,
			"lastName":  user.LastName,
			"email":     user.Email,
			"username":  user.Username,
		},
	}

	// Send JSON response
	SendJSON(w, http.StatusOK, data, "Dashboard data retrieved successfully")
}

func (app *Application) registerHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	firstname := r.FormValue("firstname")
	lastname := r.FormValue("lastname")
	email := r.FormValue("email")
	password := r.FormValue("password")

	// Validation errors map
	errors := map[string]string{}

	// Validate inputs
	if username == "" {
		errors["username"] = "Username is required"
	}
	if email == "" {
		errors["email"] = "Email is required"
	} else if !isValidEmail(email) {
		errors["email"] = "Invalid email format"
	}
	if password == "" {
		errors["password"] = "Password is required"
	} else if len(password) < 8 {
		errors["password"] = "Password must be at least 8 characters"
	}

	// If there are validation errors, return them
	if len(errors) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"data":    nil,
			"message": "Validation failed",
			"errors":  errors,
			"status":  http.StatusBadRequest,
		})
		return
	}

	// Create user in the database
	err := app.UserModel.CreateUser(r.Context(), username, firstname, lastname, email, password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"data":    nil,
			"message": "Failed to create user",
			"errors":  nil,
			"status":  http.StatusInternalServerError,
		})
		fmt.Printf("Failed to create user: %v\n", err)
		return
	}

	// Success response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data":    nil,
		"message": "User registered successfully",
		"errors":  nil,
		"status":  http.StatusCreated,
	})
}

// Helper to validate email format
func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`)
	return re.MatchString(email)
}

func (app *Application) loginHandler(w http.ResponseWriter, r *http.Request) {
	var req models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Invalid input",
			"status":  http.StatusBadRequest,
		})
		return
	}

	// Authenticate user
	user, err := app.UserModel.AuthenticateUser(r.Context(), req.Email, req.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Invalid credentials",
			"status":  http.StatusUnauthorized,
		})
		return
	}

	// Store user ID in session
	app.Session.Put(r.Context(), "userID", user.ID)

	// Send success response
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Login successful",
		"status":  http.StatusOK,
	})
}

func (app *Application) logoutHandler(w http.ResponseWriter, r *http.Request) {
	app.Session.Destroy(r.Context())

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Logged out successfully",
	})
}

func (app *Application) authStatusHandler(w http.ResponseWriter, r *http.Request) {
	userID := app.Session.GetInt(r.Context(), "userID")

	if userID == 0 {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"user": nil,
		})
		return
	}

	user, err := app.UserModel.GetUserByID(r.Context(), userID)
	if err != nil {
		http.Error(w, "Failed to get user", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"user": user,
	})
}
