package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/testgithubanjali/ai-document-summarizer/internal/config"
	"github.com/testgithubanjali/ai-document-summarizer/internal/database"
)

func main() {

	config.LoadEnv()

	log.Println("PORT:", config.GetEnv("PORT"))

	database.ConnectDB()

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "AI Document Summarizer API is running",
		})
	})

	router.Run(":" + config.GetEnv("PORT"))
}
