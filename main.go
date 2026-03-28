package main

import (
	"url-shortener/database"
	"url-shortener/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	r := gin.Default()

	r.POST("/shorten", handlers.CreateShortURL)
	r.GET("/:code", handlers.RedirectURL)

	r.Run(":8080")
}
