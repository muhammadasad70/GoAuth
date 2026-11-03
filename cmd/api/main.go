package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", healthHandler)
	fmt.Println("Server running on the port 8080")
	http.ListenAndServe(":8080", mux)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("GoAuth api is running "))
}
