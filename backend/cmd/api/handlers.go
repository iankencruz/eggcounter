package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/go-chi/chi/v5"
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

	// Fetch total eggs consumed
	totalEggs, err := app.EggModel.GetTotalEggCount(r.Context(), userID)
	if err != nil {
		SendJSON(w, http.StatusInternalServerError, nil, "Failed to fetch total egg count")
		return
	}

	// Fetch recent entries
	recentEntries, err := app.EggModel.GetRecentEggEntries(r.Context(), userID, 5)
	if err != nil {
		SendJSON(w, http.StatusInternalServerError, nil, "Failed to fetch recent egg entries")
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
		"totalEggs":     totalEggs,
		"recentEntries": recentEntries,
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
	SendJSON(w, http.StatusOK, nil, "Logged out successfully")
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

func (app *Application) getEggCountHandler(w http.ResponseWriter, r *http.Request) {
	userID := app.Session.GetInt(r.Context(), "userID")
	if userID == 0 {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	total, err := app.EggModel.GetTotalEggCount(r.Context(), userID)
	if err != nil {
		http.Error(w, "Failed to fetch egg count", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"totalEggs": total,
	})
}

func (app *Application) addEggCountHandler(w http.ResponseWriter, r *http.Request) {
	userID := app.Session.GetInt(r.Context(), "userID")
	if userID == 0 {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	type Request struct {
		Amount int `json:"amount"`
	}

	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Amount <= 0 {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err := app.EggModel.AddEggCount(r.Context(), userID, req.Amount)
	if err != nil {
		http.Error(w, "Failed to add egg count", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Egg count added successfully",
	})
}

func (app *Application) deleteEntryHandler(w http.ResponseWriter, r *http.Request) {
	userID := app.Session.GetInt(r.Context(), "userID")
	if userID == 0 {
		SendJSON(w, http.StatusUnauthorized, nil, "Unauthorized. Please log in.")
		return
	}

	// Extract entry ID from URL
	entryID := chi.URLParam(r, "id")
	if entryID == "" {
		SendJSON(w, http.StatusBadRequest, nil, "Entry ID is required")
		return
	}

	// Convert entryID to an integer
	id, err := strconv.Atoi(entryID)
	if err != nil {
		SendJSON(w, http.StatusBadRequest, nil, "Invalid entry ID")
		return
	}

	// Get the amount for this entry
	amount, err := app.EggModel.GetEntryByID(r.Context(), userID, id)
	if err != nil {
		SendJSON(w, http.StatusInternalServerError, nil, "Failed to retrieve entry amount")
		return
	}

	// Insert a new "undo" entry (insert a negative amount)
	err = app.EggModel.InsertNegativeEntry(r.Context(), userID, amount)
	if err != nil {
		SendJSON(w, http.StatusInternalServerError, nil, "Failed to add reversal entry")
		return
	}

	SendJSON(w, http.StatusOK, nil, "Entry successfully undone")
}
