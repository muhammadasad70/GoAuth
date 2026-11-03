package main

import (
	"log"
	"net/http"

	"github.com/muhammadasad70/goauth/internal/handler"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", handler.HealthHandler)
	log.Println("Server running on the port 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
