package handlers

import (
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/testgithubanjali/ai-document-summarizer/internal/dto"
	"github.com/testgithubanjali/ai-document-summarizer/internal/models"
	"github.com/testgithubanjali/ai-document-summarizer/internal/service"
	"github.com/testgithubanjali/ai-document-summarizer/internal/utils"
)

type DocumentHandler struct {
	documentService *service.DocumentService
}

func NewDocumentHandler(service *service.DocumentService) *DocumentHandler {
	return &DocumentHandler{
		documentService: service,
	}
}

func (h *DocumentHandler) Upload(c *gin.Context) {

	userIDValue, exists := c.Get("user_id")
	if !exists {
		utils.Error(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	userID, ok := userIDValue.(float64)
	if !ok {
		utils.Error(c, http.StatusUnauthorized, "Invalid user")
		return
	}

	file, err := c.FormFile("document")
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "Document is required")
		return
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))

	if ext != ".pdf" && ext != ".docx" {
		utils.Error(c, http.StatusBadRequest, "Only PDF and DOCX files are allowed")
		return
	}

	if err := os.MkdirAll("uploads", os.ModePerm); err != nil {
		utils.Error(c, http.StatusInternalServerError, "Failed to create upload directory")
		return
	}

	fileName := strconv.FormatInt(time.Now().UnixNano(), 10) + ext
	filePath := filepath.Join("uploads", fileName)

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		utils.Error(c, http.StatusInternalServerError, "Failed to save file")
		return
	}

	document := models.Document{
		UserID:   uint(userID),
		FileName: file.Filename,
		FilePath: filePath,
		FileType: ext,
	}

	summary, err := h.documentService.ProcessDocument(&document)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response := dto.ProcessDocumentResponse{
		ID:       document.ID,
		FileName: document.FileName,
		FileType: document.FileType,
		Summary:  summary,
		Message:  "Document processed successfully",
	}

	utils.Success(c, http.StatusCreated, response)
}
