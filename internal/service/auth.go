package service

import (
	"context"
	"errors"
	"fmt"
	"net/mail"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/muhammadasad70/goauth/internal/model"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidEmail       = errors.New("invalid email")
	ErrInvalidPassword    = errors.New("password must contain atleast 8 character")
	ErrPasswordTooLong    = errors.New("password exceeds bcrypt limit")
	ErrInvalidCredentials = errors.New("invalid email and password")
)

type UserRepository interface {
	CreateUser(
		ctx context.Context,
		email string,
		passwordHash string,

	) (int64, error)

	GetUserByEmail(
		ctx context.Context,
		email string,

	) (model.User, error)
}
type AuthService struct {
	repo UserRepository
}

func NewAuthService(repo UserRepository) *AuthService {
	return &AuthService{
		repo: repo,
	}

}
func (s *AuthService) RegisterUser(
	ctx context.Context,
	email string,
	passwordHash string,

) (int64, error) {
	email = strings.TrimSpace(email)

	addr, err := mail.ParseAddress(email)
	if err != nil || addr.Address != email {
		return 0, ErrInvalidEmail

	}

	if len(passwordHash) < 8 {
		return 0, ErrInvalidPassword
	}

	if len([]byte(passwordHash)) > 78 {
		return 0, ErrPasswordTooLong

	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(passwordHash),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return 0, fmt.Errorf("hash password :%w", err)
	}

	userID, err := s.repo.CreateUser(
		ctx,
		email,
		string(hashedPassword),
	)
	if err != nil {
		return 0, fmt.Errorf("register user : %w", err)
	}
	return userID, nil
}
func (s *AuthService) LoginUser(
	ctx context.Context,
	email string,
	password string,
) (int64, error) {
	email = strings.TrimSpace(email)

	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, ErrInvalidCredentials
		}
		return 0, fmt.Errorf("login user: %w", err)
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(password),
	)
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return 0, ErrInvalidCredentials
		}
		return 0, fmt.Errorf("verify user: %w", err)
	}
	return user.ID, nil

}
