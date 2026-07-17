package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/testgithubanjali/ai-document-summarizer/internal/dto"
	"github.com/testgithubanjali/ai-document-summarizer/internal/models"
	"github.com/testgithubanjali/ai-document-summarizer/internal/service"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) Register(c *gin.Context) {

	var req dto.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	err := h.userService.Register(&user)
	if err != nil {

		if errors.Is(err, service.ErrEmailAlreadyExists) {
			c.JSON(http.StatusConflict, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}

	response := dto.RegisterResponse{
		ID:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		Message: "User registered successfully",
	}

	c.JSON(http.StatusCreated, response)
}
