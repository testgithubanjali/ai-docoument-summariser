package service

import (
	"github.com/testgithubanjali/ai-document-summarizer/internal/models"
	"github.com/testgithubanjali/ai-document-summarizer/internal/repository"
)

type DocumentService struct {
	documentRepo *repository.DocumentRepository
}

func NewDocumentService(repo *repository.DocumentRepository) *DocumentService {
	return &DocumentService{
		documentRepo: repo,
	}
}

func (s *DocumentService) Create(document *models.Document) error {
	return s.documentRepo.Create(document)
}
