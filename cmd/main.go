package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/testgithubanjali/ai-document-summarizer/internal/config"
	"github.com/testgithubanjali/ai-document-summarizer/internal/database"
	"github.com/testgithubanjali/ai-document-summarizer/internal/routes"
)

func main() {

	config.LoadEnv()

	log.Println("PORT:", config.GetEnv("PORT"))

	database.ConnectDB()

	router := gin.Default()

	routes.SetupRoutes(router)

	err := router.Run(":" + config.GetEnv("PORT"))
	if err != nil {
		log.Fatal(err)
	}
}
