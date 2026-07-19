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

	// User dependencies
	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	// Document dependencies
	documentRepo := repository.NewDocumentRepository()
	summaryRepo := repository.NewSummaryRepository()

	documentService := service.NewDocumentService(
		documentRepo,
		summaryRepo,
	)
	documentHandler := handlers.NewDocumentHandler(documentService)

	// Summary dependencies
	summaryService := service.NewSummaryService(summaryRepo)
	summaryHandler := handlers.NewSummaryHandler(summaryService)

	// Public routes
	router.POST("/register", userHandler.Register)
	router.POST("/login", userHandler.Login)

	// Protected routes
	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())

	// User
	protected.GET("/profile", userHandler.GetProfile)

	// Document
	protected.POST("/documents/upload", documentHandler.Upload)

	// Summary
	protected.GET("/summaries", summaryHandler.GetAll)
	protected.GET("/summaries/:id", summaryHandler.GetByID)
	protected.DELETE("/summaries/:id", summaryHandler.Delete)
}
