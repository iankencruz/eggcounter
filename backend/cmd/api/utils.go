package main

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Data    interface{} `json:"data"` // Correctly wraps "data"
	Message string      `json:"message"`
	Status  int         `json:"status"`
}

func SendJSON(w http.ResponseWriter, statusCode int, data interface{}, message string) {
	response := APIResponse{
		Data:    data, // No extra wrapping
		Message: message,
		Status:  statusCode,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
