package handlers

import (
	"url-shortener/database"
	"url-shortener/utils"

	"github.com/gin-gonic/gin"
)

func CreateShortURL(c *gin.Context) {
	var body struct {
		URL string `json:"url"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	code := utils.GenerateCode(6)

	_, err := database.DB.Exec(
		"INSERT INTO urls (short_code, original_url) VALUES ($1, $2)",
		code, body.URL,
	)

	if err != nil {
		c.JSON(500, gin.H{"error": "DB error"})
		return
	}
        host := c.Request.Host
        shortURL := "http://" + host + "/" + code
        c.JSON(200, gin.H{
    	   "short_url": shortURL,
        })
} 

func RedirectURL(c *gin.Context) {
	code := c.Param("code")

	var original string

	err := database.DB.QueryRow(
		"SELECT original_url FROM urls WHERE short_code=$1",
		code,
	).Scan(&original)

	if err != nil {
		c.JSON(404, gin.H{"error": "Not found"})
		return
	}

	database.DB.Exec(
		"UPDATE urls SET clicks = clicks + 1 WHERE short_code=$1",
		code,
	)

	c.Redirect(302, original)
}
