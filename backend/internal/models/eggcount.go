package models

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// EggCount represents an egg consumption record.
type EggCount struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

// EggModel handles database operations for the eggcount table.
type EggModel struct {
	DB *pgxpool.Pool
}

// NewEggModel creates a new instance of EggModel.
func NewEggModel(db *pgxpool.Pool) *EggModel {
	return &EggModel{DB: db}
}

// AddEggCount adds a new egg consumption record for the user.
func (m *EggModel) AddEggCount(ctx context.Context, userID, amount int) error {
	query := `
		INSERT INTO eggcount (user_id, amount)
		VALUES ($1, $2)
	`
	_, err := m.DB.Exec(ctx, query, userID, amount)
	return err
}

// GetTotalEggCount retrieves the total number of eggs consumed by a user.
func (m *EggModel) GetTotalEggCount(ctx context.Context, userID int) (int, error) {
	var totalEggs int
	query := `
		SELECT COALESCE(SUM(amount), 0)
		FROM eggcount
		WHERE user_id = $1
	`
	err := m.DB.QueryRow(ctx, query, userID).Scan(&totalEggs)
	return totalEggs, err
}

// GetRecentEggEntries retrieves the most recent egg consumption records for a user.
func (m *EggModel) GetRecentEggEntries(ctx context.Context, userID int, limit int) ([]EggCount, error) {
	query := `
		SELECT id, user_id, amount, created_at
		FROM eggcount
		WHERE user_id = $1
		ORDER BY created_at DESC
		LIMIT $2
	`

	rows, err := m.DB.Query(ctx, query, userID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entries []EggCount
	for rows.Next() {
		var eggCount EggCount
		if err := rows.Scan(&eggCount.ID, &eggCount.UserID, &eggCount.Amount, &eggCount.CreatedAt); err != nil {
			return nil, err
		}
		entries = append(entries, eggCount)
	}

	return entries, nil
}
