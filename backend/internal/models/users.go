package models

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int
	Username  string `json:"username"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"-"`
	CreatedAt string `json:"created_at"`
}

type UserModel struct {
	DB *pgxpool.Pool
}

func NewUserModel(db *pgxpool.Pool) *UserModel {
	return &UserModel{DB: db}
}

func (m *UserModel) CreateUser(ctx context.Context, username, firstName, lastName, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Print("hashing error")
		return err
	}

	query := `
	INSERT INTO users (username, email, first_name, last_name, password_hash) 
	VALUES ($1, $2, $3, $4, $5)`

	_, err = m.DB.Exec(ctx, query, username, email, firstName, lastName, string(hashedPassword))

	return err
}

func (m *UserModel) AuthenticateUser(ctx context.Context, email, password string) (*User, error) {
	var user User

	// Use positional placeholders ($1)
	query := `
    SELECT id, username, email, first_name, last_name, password_hash 
    FROM users 
    WHERE email = $1`

	// Execute the query with email as the parameter
	row := m.DB.QueryRow(ctx, query, email)

	// Scan the results into the user struct
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.FirstName, &user.LastName, &user.Password)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Compare the hashed password from DB with the user-provided password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return &user, nil
}

func (m *UserModel) GetUserByID(ctx context.Context, userID int) (*User, error) {
	var user User

	query := `
		SELECT id, username, email, first_name, last_name 
		FROM users 
		WHERE id = $1
	`

	err := m.DB.QueryRow(ctx, query, userID).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.FirstName,
		&user.LastName,
	)

	if err != nil {
		return nil, errors.New("user not found")
	}

	return &user, nil
}
