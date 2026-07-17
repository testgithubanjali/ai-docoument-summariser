package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, status int, data interface{}) {
	c.JSON(status, data)
}

func Error(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{
		"error": message,
	})
}

func ValidationError(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error": err.Error(),
	})
}
