package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/testgithubanjali/ai-document-summarizer/internal/service"
)

type SummaryHandler struct {
	service *service.SummaryService
}

func NewSummaryHandler(service *service.SummaryService) *SummaryHandler {
	return &SummaryHandler{
		service: service,
	}
}

func (h *SummaryHandler) GetAll(c *gin.Context) {

	userID := c.GetUint("userID")

	summaries, err := h.service.GetAll(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, summaries)
}

func (h *SummaryHandler) GetByID(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	summary, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Summary not found",
		})
		return
	}

	c.JSON(http.StatusOK, summary)
}

func (h *SummaryHandler) Delete(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	if err := h.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Summary deleted successfully",
	})
}
