package service

import (
	"errors"
	"os"

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

	var (
		text string
		err  error
	)

	switch document.FileType {

	case ".pdf":
		text, err = utils.ExtractPDFText(document.FilePath)
		if err != nil {
			_ = os.Remove(document.FilePath)
			return "", err
		}

	case ".docx":
		_ = os.Remove(document.FilePath)
		return "", errors.New("DOCX processing is not implemented yet")

	default:
		_ = os.Remove(document.FilePath)
		return "", errors.New("unsupported file type")
	}

	summary, err := ai.GenerateSummary(text)
	if err != nil {
		_ = os.Remove(document.FilePath)
		return "", err
	}

	if err := s.documentRepo.Create(document); err != nil {
		_ = os.Remove(document.FilePath)
		return "", err
	}

	return summary, nil
}
