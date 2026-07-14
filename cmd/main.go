package main

import (
	"github.com/gin-gonic/gin"

	"github.com/yourusername/ai-document-summarizer/internal/config"
)

func main() {
	config.LoadEnv()

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "AI Document Summarizer API is running",
		})
	})

	router.Run(":" + config.GetEnv("PORT"))
}
