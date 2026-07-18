package repository

import (
	"github.com/testgithubanjali/ai-document-summarizer/internal/database"
	"github.com/testgithubanjali/ai-document-summarizer/internal/models"
)

type DocumentRepository struct{}

func NewDocumentRepository() *DocumentRepository {
	return &DocumentRepository{}
}

func (r *DocumentRepository) Create(document *models.Document) error {
	return database.DB.Create(document).Error
}
