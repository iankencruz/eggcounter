package models

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserModel struct {
	DB           *pgxpool.Pool
	SessionStore *sessions.CookieStore
}

// WithDB adds the database connection to the context
func WithDB(ctx context.Context, db *pgxpool.Pool) context.Context {
	return context.WithValue(ctx, "db", db)
}

// GetDB retrieves the database connection from the context
func GetDB(ctx context.Context) (*pgxpool.Pool, error) {
	db, ok := ctx.Value("db").(*pgxpool.Pool)
	if !ok {
		return nil, errors.New("could not retrieve database connection from context")
	}
	return db, nil
}

// AuthenticateUser authenticates a user by username and password
func (um *UserModel) AuthenticateUser(ctx context.Context, w http.ResponseWriter, r *http.Request, username, password string) (*User, error) {
	var user User
	query := `SELECT id, username, password, email FROM users WHERE username = :username`
	args := pgx.NamedArgs{"username": username}
	row := um.DB.QueryRow(ctx, query, args)
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	if err != nil {
		return nil, errors.New("invalid username or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid username or password")
	}

	session, _ := um.SessionStore.Get(r, "user-session")
	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7, // 7 days
		HttpOnly: true,
	}
	session.Values["user_id"] = user.ID

	err = session.Save(r, w)
	if err != nil {
		log.Println("Failed to save session:", err)
		return nil, errors.New("failed to save session")
	}

	return &user, nil
}

// GetUserByID retrieves a user by their ID
func (um *UserModel) GetUserByID(ctx context.Context, userID int) (*User, error) {
	var user User
	query := `SELECT id, username, email FROM users WHERE id = :id`
	args := pgx.NamedArgs{"id": userID}
	row := um.DB.QueryRow(ctx, query, args)
	err := row.Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// CreateUser inserts a new user into the database
func (um *UserModel) CreateUser(ctx context.Context, username, password, email string) (*User, error) {
	// Hash the password before storing it in the database
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	var user User
	query := `INSERT INTO users (username, password, email) VALUES (:username, :password, :email) RETURNING id, username, email`
	args := pgx.NamedArgs{"username": username, "password": string(hashedPassword), "email": email}
	row := um.DB.QueryRow(ctx, query, args)
	err = row.Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (um *UserModel) GetSessionUserID(r *http.Request) (int, error) {
	if um.SessionStore == nil {
		return 0, errors.New("session store is not initialized")
	}

	session, err := um.SessionStore.Get(r, "user-session")
	if err != nil {
		return 0, errors.New("failed to get session: " + err.Error())
	}

	// Check if session is nil
	if session == nil {
		return 0, errors.New("session is nil")
	}

	// Check if session values exist
	if session.Values == nil {
		return 0, errors.New("session values are nil")
	}

	// Check if user_id exists in session
	userID, ok := session.Values["user_id"]
	if !ok {
		return 0, errors.New("user ID not found in session")
	}

	// Check if the user_id is an int
	userIDInt, ok := userID.(int)
	if !ok {
		return 0, errors.New("user ID is not of type int")
	}

	return userIDInt, nil
}
