package main

import (
	"github.com/gin-gonic/gin"

	"github.com/manmithsm/url-shortener/config"
	"github.com/manmithsm/url-shortener/internal/handlers"
	"github.com/manmithsm/url-shortener/internal/repository"
	"github.com/manmithsm/url-shortener/internal/services"
)

func main() {
	// DB connection
	db := config.ConnectDB()
	defer db.Close()

	// Dependency injection
	repo := repository.NewURLRepository(db)
	service := services.NewURLService(repo)
	handler := handlers.NewURLHandler(service)

	// Router
	r := gin.Default()

	// Health
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "OK"})
	})

	// API
	r.POST("/api/v1/shorten", handler.Shorten)
	r.GET("/:shortCode", handler.Redirect)

	// Run server
	r.Run(":8080")
}
