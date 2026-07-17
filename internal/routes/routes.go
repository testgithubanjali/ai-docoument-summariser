package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/testgithubanjali/ai-document-summarizer/internal/handlers"
	"github.com/testgithubanjali/ai-document-summarizer/internal/repository"
	"github.com/testgithubanjali/ai-document-summarizer/internal/service"
)

func SetupRoutes(router *gin.Engine) {

	// Health Check
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "AI Document Summarizer API is running",
		})
	})

	// Dependency Injection
	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	// User Routes
	router.POST("/register", userHandler.Register)
}
