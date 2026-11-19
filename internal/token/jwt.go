package token

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateAccessToken(userID int64, secret []byte) (string, error) {
	if len(secret) < 32 {
		return "", errors.New("Jwt secret must be atleast 23 bytes")
	}

	now := time.Now()
	claims := jwt.RegisteredClaims{
		Subject:   strconv.FormatInt(userID, 10),
		Issuer:    "goauth",
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(15 * time.Minute)),
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString(secret)
}
func ValidateAccessToken(tokenString string, secret []byte) (int64, error) {
	claims := &jwt.RegisteredClaims{}
	parsedToken, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(t *jwt.Token) (interface{}, error) {
			return secret, nil

		},
		jwt.WithValidMethods([]string{"HS256"}),
		jwt.WithExpirationRequired(),
		jwt.WithIssuer("goauth"),
	)
	if err != nil {
		return 0, fmt.Errorf("validate token:%w", err)
	}
	if !parsedToken.Valid {
		return 0, errors.New("invalid token")
	}
	userID, err := strconv.ParseInt(claims.Subject, 10, 64)
	if err != nil || userID <= 0 {
		return 0, errors.New("invalid user ID")

	}
	return userID, nil
}
