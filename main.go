package main

import (
	"github.com/gin-gonic/gin"

	"os"
)

func main() {
	r := gin.Defeault()

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	r.Get("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "this is the first APi",
		})
	})
}
