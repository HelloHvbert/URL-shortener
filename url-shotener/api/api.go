package api

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"example.com/db"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Redirect to the original URL
func redirect(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	isRedirected := id[0] != '!'

	if !isRedirected {
		id = id[1:]
		str, _ := url.PathUnescape(id)
		originalURL, err := db.GetOriginalUrl(str)

		if err != nil {
			// c.JSON(404, gin.H{"error": "URL not founddd"})
			c.HTML(404, "index.tmpl", gin.H{"message": "URL not found"},)
			return
		}
	
		c.HTML(200, "index.tmpl", gin.H{"site": originalURL, "message": "Go to"},)
		return
	}

	originalURL, err := db.GetOriginalUrl(id)

	if err != nil {
		log.Printf("Error getting original URL for id %s: %v", id, err)
		c.JSON(404, gin.H{"error": "URL not found"})
		return
	}
	
	c.Redirect(301, originalURL) // Ensure this is the correct URL
}

// Initialize the API
func Run() {
	router := gin.Default()
	router.LoadHTMLGlob("api/templates/*")
	router.GET("/:id", redirect)

	// Get PORT from .env
	err := godotenv.Load(".env")
	if err != nil {
			log.Fatalf("Error loading .env file: %s", err)
	}
	PORT := os.Getenv("API_PORT")

	fmt.Printf("Server running at http://localhost:%s\n", PORT)
	router.Run(fmt.Sprintf(":%s", PORT))
}