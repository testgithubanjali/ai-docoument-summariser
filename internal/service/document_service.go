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
	summaryRepo  *repository.SummaryRepository
}

func NewDocumentService(
	documentRepo *repository.DocumentRepository,
	summaryRepo *repository.SummaryRepository,
) *DocumentService {
	return &DocumentService{
		documentRepo: documentRepo,
		summaryRepo:  summaryRepo,
	}
}

func (s *DocumentService) ProcessDocument(document *models.Document) (string, error) {
	var (
		text    string
		summary string
		err     error
	)

	// Extract text based on file type
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

	// Generate summary using Gemini
	summary, err = ai.GenerateSummary(text)
	if err != nil {
		_ = os.Remove(document.FilePath)
		return "", err
	}

	// Save document metadata
	if err := s.documentRepo.Create(document); err != nil {
		_ = os.Remove(document.FilePath)
		return "", err
	}

	// Save summary
	summaryModel := &models.Summary{
		DocumentID: document.ID,
		Content:    summary,
		ModelUsed:  "gemini-2.5-flash",
	}

	if err := s.summaryRepo.Create(summaryModel); err != nil {
		_ = os.Remove(document.FilePath)
		return "", err
	}

	return summary, nil
}
