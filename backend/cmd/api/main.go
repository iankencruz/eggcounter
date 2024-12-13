package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/iankencruz/eggcounter/backend/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

type application struct {
	DB           *pgxpool.Pool
	User         *models.UserModel
	SessionStore *sessions.CookieStore
}

func main() {
	// Initialize session store

	curDir, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	// Load environment variables from the .env file
	err = godotenv.Load(curDir + "/backend/.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Create database connection string
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	sessionStore := sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Initialize database connection pool
	pool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer pool.Close()

	userModel := &models.UserModel{
		SessionStore: sessionStore,
	}

	// Create application struct
	app := &application{
		DB:           pool,
		User:         userModel,
		SessionStore: sessionStore,
	}

	// Start the server on port 5050
	log.Println("Starting GO API server on :3000")
	if err := http.ListenAndServe(":3000", app.routes()); err != nil {
		log.Fatal(err)
	}
}
