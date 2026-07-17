package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/testgithubanjali/ai-document-summarizer/internal/dto"
	"github.com/testgithubanjali/ai-document-summarizer/internal/models"
	"github.com/testgithubanjali/ai-document-summarizer/internal/service"
	"github.com/testgithubanjali/ai-document-summarizer/internal/utils"
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
		utils.ValidationError(c, err)
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
			utils.Error(c, http.StatusConflict, err.Error())
			return
		}

		utils.Error(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	response := dto.RegisterResponse{
		ID:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		Message: "User registered successfully",
	}

	utils.Success(c, http.StatusCreated, response)
}
func (h *UserHandler) Login(c *gin.Context) {

	var req dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationError(c, err)
		return
	}

	user, err := h.userService.Login(req.Email, req.Password)

	if err != nil {

		if errors.Is(err, service.ErrInvalidCredentials) {
			utils.Error(c, http.StatusUnauthorized, err.Error())
			return
		}

		utils.Error(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	response := dto.LoginResponse{
		ID:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		Message: "Login successful",
	}

	utils.Success(c, http.StatusOK, response)
}
