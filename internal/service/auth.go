package service

import (
	"context"
	"errors"
	"fmt"
	"net/mail"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidEmail    = errors.New("invalid email")
	ErrInvalidPassword = errors.New("password must contain atleast 8 character")
	ErrPasswordTooLong = errors.New("password exceeds bcrypt limit")
)

type UserRepository interface {
	CreateUser(
		ctx context.Context,
		email string,
		passwordHash string,

	) (int64, error)
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
