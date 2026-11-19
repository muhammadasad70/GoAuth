package handler

import (
	"encoding/json"
	"net/http"

	"github.com/muhammadasad70/goauth/internal/middleware"
)

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := middleware.GetUserID(r.Context())
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	response := map[string]any{
		"user_id": userID,
		"message": "welcome to your profile",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
