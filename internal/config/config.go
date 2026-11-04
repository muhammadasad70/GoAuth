package config

import (
	"os"
)

// config struct groups application configuration at one place
type Config struct {
	Port string
}

func Load() Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return Config{
		Port: port,
	}
}
