package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/muhammadasad70/goauth/internal/model"
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
		VALUES($1,$2)
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
func (r *Userrepository) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	query := `
		SELECT id,email,password_hash,created_at
		FROM public.users
		WHERE email=$1
	`
	var user model.User

	err := r.db.QueryRow(ctx, query, email).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,

		&user.CreatedAt,
	)
	if err != nil {
		return model.User{}, fmt.Errorf("get user by email: %w", err)
	}
	return user, nil
}
