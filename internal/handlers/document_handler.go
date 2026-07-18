package handlers

import (
	"log"
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

	userID := uint(userIDValue.(float64))

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

	// Create uploads directory if it doesn't exist
	if err := os.MkdirAll("uploads", os.ModePerm); err != nil {
		utils.Error(c, http.StatusInternalServerError, "Failed to create upload directory")
		return
	}

	// Generate unique filename
	fileName := strconv.FormatInt(time.Now().UnixNano(), 10) + ext
	filePath := filepath.Join("uploads", fileName)

	// Save uploaded file
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		utils.Error(c, http.StatusInternalServerError, "Failed to save file")
		return
	}

	// Extract text only from PDF files (temporary testing)
	var extractedText string

	if ext == ".pdf" {
		extractedText, err = utils.ExtractPDFText(filePath)
		if err != nil {
			utils.Error(c, http.StatusInternalServerError, "Failed to extract PDF text")
			return
		}

		log.Println("========== EXTRACTED PDF TEXT ==========")
		log.Println(extractedText)
		log.Println("========================================")
	}

	document := models.Document{
		UserID:   userID,
		FileName: file.Filename,
		FilePath: filePath,
		FileType: ext,
	}

	if err := h.documentService.Create(&document); err != nil {
		utils.Error(c, http.StatusInternalServerError, "Failed to save document")
		return
	}

	response := dto.UploadResponse{
		ID:       document.ID,
		FileName: document.FileName,
		FileType: document.FileType,
		Message:  "Document uploaded successfully",
	}

	utils.Success(c, http.StatusCreated, response)
}
