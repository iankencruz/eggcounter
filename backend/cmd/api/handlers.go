package main

import (
	"fmt"
	"net/http"
)

func (app *application) DashboardHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Dashboard API")
	w.Write([]byte("Dashboard API"))
}
