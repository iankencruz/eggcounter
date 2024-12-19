package models

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Friend struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type FriendModel struct {
	DB *pgxpool.Pool
}

// AddFriendRequest adds a friend request to the database
// SendFriendRequest adds a friend request to the database
func (m *FriendModel) SendFriendRequest(ctx context.Context, fromUserID, toUserID int) error {
	query := `
		INSERT INTO friends (from_user_id, to_user_id, status, created_at) 
		VALUES ($1, $2, 'pending', NOW())
	`
	_, err := m.DB.Exec(ctx, query, fromUserID, toUserID)
	return err
}

// eggcounter/backend/internal/models/friends.go

// File: internal/models/friends.go
func (m *FriendModel) GetFriendsList(ctx context.Context, userID int) ([]User, error) {
	query := `
		SELECT 
			users.id, 
			users.username, 
			users.email, 
			friends.created_at 
		FROM friends 
		JOIN users ON friends.friend_id = users.id 
		WHERE friends.user_id = $1 
		  AND friends.status = 'accepted'
	`
	rows, err := m.DB.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var friends []User
	for rows.Next() {
		var friend User
		err := rows.Scan(&friend.ID, &friend.Username, &friend.Email, &friend.CreatedAt)
		if err != nil {
			return nil, err
		}
		friends = append(friends, friend)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return friends, nil
}

// eggcounter/backend/internal/models/friends.go

func (m *FriendModel) AcceptFriendRequest(ctx context.Context, friendID string) error {
	query := `
		UPDATE friends 
		SET status = 'accepted' 
		WHERE id = $1
	`
	_, err := m.DB.Exec(ctx, query, friendID)
	return err
}

func (m *FriendModel) RejectFriendRequest(ctx context.Context, friendID string) error {
	query := `
		UPDATE friends 
		SET status = 'rejected' 
		WHERE id = $1
	`
	_, err := m.DB.Exec(ctx, query, friendID)
	return err
}

func (m *FriendModel) GetFriendRequests(ctx context.Context, userID int) ([]User, error) {
	query := `
		SELECT 
			users.id, 
			users.username, 
			users.email, 
			friends.created_at 
		FROM friends 
		JOIN users ON friends.user_id = users.id 
		WHERE friends.friend_id = $1 
		  AND friends.status = 'pending'
	`

	rows, err := m.DB.Query(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("error querying friend requests: %v", err)
	}
	defer rows.Close()

	var friendRequests []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("error scanning friend request: %v", err)
		}
		friendRequests = append(friendRequests, user)
	}

	return friendRequests, nil
}
