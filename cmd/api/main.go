package main

import (
	"log"
	"net/http"

	"github.com/muhammadasad70/goauth/internal/config"
	"github.com/muhammadasad70/goauth/internal/handler"
)

func main() {
	cfg := config.Load()
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", handler.HealthHandler)
	log.Println("Server running on the port", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, mux))
}
