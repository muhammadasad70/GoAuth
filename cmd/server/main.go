package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/muhammadasad70/go-auth-service/internal/config"
	"github.com/muhammadasad70/go-auth-service/internal/database"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load config", err)
	}
	db, err := database.NewPostgres(cfg.DatabaseUrl)
	if err != nil {
		log.Fatal("failed to connect postgres", err)
	}
	defer db.Close()
	redisClient, err := database.NewRedis(
		cfg.RedisAdrr,
		cfg.RedisPassword,
	)
	if err != nil {
		log.Fatal("failed to connect redis", err)
	}
	defer redisClient.Close()
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	log.Printf("Server is running on the port %s", cfg.ServerPort)

	if err := router.Run(":" + cfg.ServerPort); err != nil {
		log.Fatal("Failed to start server:", err)
	}

}
