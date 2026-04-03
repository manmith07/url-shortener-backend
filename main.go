package main
import "os"
import (
	"url-shortener/database"
	"url-shortener/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	r := gin.Default()

	// ✅ Strong CORS config
	r.Use(cors.New(cors.Config{
    AllowOrigins: []string{
        "http://localhost:3000",
        "https://v0-nanolink.vercel.app",
    },
    AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowHeaders: []string{"Origin", "Content-Type", "Accept"},
    AllowCredentials: true,
}))

	// ✅ VERY IMPORTANT: handle OPTIONS explicitly
	r.OPTIONS("/*path", func(c *gin.Context) {
		c.Status(204)
	})

	// Routes
	r.POST("/shorten", handlers.CreateShortURL)
	r.GET("/:code", handlers.RedirectURL)

	port := os.Getenv("Port")
        if port ==""{
	   port = "8080"
        } 
        r.Run(":" + port)
        	
}
