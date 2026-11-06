package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/muhammadasad70/goauth/internal/config"
	"github.com/muhammadasad70/goauth/internal/database"
	"github.com/muhammadasad70/goauth/internal/handler"
)

func main() {
	cfg := config.Load()

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL is required")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := database.Connect(ctx, databaseURL)
	if err != nil {
		log.Fatal("Database connection failed", err)
	}
	defer pool.Close()
	log.Println("Successfully connected to PostgreSQL")

	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", handler.HealthHandler)

	log.Println("Server running on the port", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, mux))
}
