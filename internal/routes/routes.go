package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/testgithubanjali/ai-document-summarizer/internal/handlers"
	"github.com/testgithubanjali/ai-document-summarizer/internal/middleware"
	"github.com/testgithubanjali/ai-document-summarizer/internal/repository"
	"github.com/testgithubanjali/ai-document-summarizer/internal/service"
)

func SetupRoutes(router *gin.Engine) {

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "AI Document Summarizer API is running",
		})
	})

	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	router.POST("/register", userHandler.Register)
	router.POST("/login", userHandler.Login)

	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())

	protected.GET("/profile", userHandler.GetProfile)
}
