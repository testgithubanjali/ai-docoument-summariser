package service

import (
	"errors"

	"github.com/testgithubanjali/ai-document-summarizer/internal/ai"
	"github.com/testgithubanjali/ai-document-summarizer/internal/models"
	"github.com/testgithubanjali/ai-document-summarizer/internal/repository"
	"github.com/testgithubanjali/ai-document-summarizer/internal/utils"
)

type DocumentService struct {
	documentRepo *repository.DocumentRepository
}

func NewDocumentService(repo *repository.DocumentRepository) *DocumentService {
	return &DocumentService{
		documentRepo: repo,
	}
}

func (s *DocumentService) ProcessDocument(document *models.Document) (string, error) {

	var text string
	var err error

	switch document.FileType {

	case ".pdf":
		text, err = utils.ExtractPDFText(document.FilePath)
		if err != nil {
			return "", err
		}

	case ".docx":
		return "", errors.New("DOCX processing is not implemented yet")

	default:
		return "", errors.New("unsupported file type")
	}

	summary, err := ai.GenerateSummary(text)
	if err != nil {
		return "", err
	}

	if err := s.documentRepo.Create(document); err != nil {
		return "", err
	}

	return summary, nil
}
