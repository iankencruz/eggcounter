package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/pgxstore"
	"github.com/alexedwards/scs/v2"
	"github.com/iankencruz/eggcounter/backend/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

type Application struct {
	DB        *pgxpool.Pool
	Session   *scs.SessionManager
	UserModel *models.UserModel
}

func main() {

	curDir, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	// Load environment variables from the .env file
	err = godotenv.Load(curDir + "/backend/.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dsn := os.Getenv("DATABASE_URL")
	dbpool, err := pgxpool.New(context.Background(), dsn)
	log.Printf("Connecting to database: %s", os.Getenv("DATABASE_URL"))

	if err != nil {
		log.Fatal("Unable to connect to database: ", err)
	}
	defer dbpool.Close()

	// Use PingContext to check the connection
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := dbpool.Ping(ctx); err != nil {
		log.Fatalf("Database ping failed: %v", err)
	}

	// Initialize PostgreSQL-backed session manager
	sessionManager := scs.New()
	sessionManager.Store = pgxstore.New(dbpool)
	sessionManager.Lifetime = 24 * time.Hour // Set session lifetime
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode
	sessionManager.Cookie.Secure = false // Set to true in production

	app := &Application{

		DB:        dbpool,
		Session:   sessionManager,
		UserModel: &models.UserModel{DB: dbpool},
	}

	// 3. Start the server
	srv := &http.Server{
		Addr:    ":8080",
		Handler: app.routes(),
	}

	log.Println("Starting server on :8080")

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal("Server failed to start: ", err)
	}

}
