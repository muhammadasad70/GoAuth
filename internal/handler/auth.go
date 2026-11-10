package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/muhammadasad70/goauth/internal/service"
)

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	UserID  int64  `json:"user_id"`
	Message string `json:"message"`
}

type AuthService interface {
	RegisterUser(
		ctx context.Context,
		email string,
		password string,
	) (int64, error)
}
type AuthHandler struct {
	service AuthService
}

func NewAuthHandler(service AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid Body", http.StatusBadRequest)
		return
	}

	userID, err := h.service.RegisterUser(
		r.Context(),
		req.Email,
		req.Password,
	)
	if err != nil {
		if errors.Is(err, service.ErrInvalidEmail) ||
			errors.Is(err, service.ErrInvalidPassword) ||
			errors.Is(err, service.ErrInvalidPassword) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	response := RegisterResponse{
		UserID:  userID,
		Message: "User register successfully",
	}
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
