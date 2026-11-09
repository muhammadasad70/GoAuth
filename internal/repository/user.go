package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Userrepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *Userrepository {
	return &Userrepository{
		db: db,
	}
}

func (r *Userrepository) CreateUser(ctx context.Context, email string, passwordHash string) (int64, error) {
	query := `
		INSERT INTO users(email,password_hash)
		VALUE($1,$2)
		RETURNING Id
	`
	var userID int64
	err := r.db.QueryRow(
		ctx,
		query,
		email,
		passwordHash,
	).Scan(&userID)
	if err != nil {
		return 0, fmt.Errorf("create user: %w", err)
	}
	return userID, nil

}
