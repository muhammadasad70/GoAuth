package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/muhammadasad70/goauth/internal/token"
)

type contextKey string

const userIDKey contextKey = "userID"

func AuthMiddleware(secret []byte) func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			authHeader := r.Header.Get("Authorization")

			parts := strings.Fields(authHeader)

			if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			userID, err := token.ValidateAccessToken(parts[1], secret)

			if err != nil {
				http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), userIDKey, userID)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
func GetUserID(ctx context.Context) (int64, bool) {
	userID, ok := ctx.Value(userIDKey).(int64)
	return userID, ok
}
