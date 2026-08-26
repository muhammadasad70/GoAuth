package config

import (
	"fmt"
	"os"
)

type Config struct {
	ServerPort    string
	DatabaseUrl   string
	RedisAdrr     string
	RedisPassword string
	JwtSecrect    string
}

func Load() (*Config, error) {
	cfg := &Config{
		ServerPort:    getEnv("Server_Port", "8080"),
		DatabaseUrl:   os.Getenv("Database_Url"),
		RedisAdrr:     getEnv("Redis_Address", "Localhost:6379"),
		RedisPassword: os.Getenv("Redis_Password"),
		JwtSecrect:    os.Getenv("JWt_Secret"),
	}

	if cfg.DatabaseUrl == "" {
		return nil, fmt.Errorf("Database_Url Required")

	}
	if cfg.JwtSecrect == "" {
		return nil, fmt.Errorf("Jwt_Secret Required")
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)

	if value == "" {
		return fallback
	}
	return value
}
