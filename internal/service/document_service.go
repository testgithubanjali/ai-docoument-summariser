package service

import (
	"errors"
	"log"
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

	log.Println("========== Document Processing Started ==========")
	log.Println("File Name:", document.FileName)
	log.Println("File Path:", document.FilePath)
	log.Println("File Type:", document.FileType)

	// Extract text based on file type
	switch document.FileType {

	case ".pdf":
		log.Println("Extracting text from PDF...")

		text, err = utils.ExtractPDFText(document.FilePath)
		if err != nil {
			log.Println("PDF Extraction Error:", err)
			_ = os.Remove(document.FilePath)
			return "", err
		}

		log.Println("PDF text extracted successfully.")
		log.Printf("Extracted %d characters\n", len(text))

	case ".docx":
		_ = os.Remove(document.FilePath)
		return "", errors.New("DOCX processing is not implemented yet")

	default:
		_ = os.Remove(document.FilePath)
		return "", errors.New("unsupported file type")
	}

	// Generate summary
	log.Println("Generating summary using Gemini...")

	summary, err = ai.GenerateSummary(text)
	if err != nil {
		log.Println("Gemini Error:", err)
		_ = os.Remove(document.FilePath)
		return "", err
	}

	log.Println("Summary generated successfully.")
	log.Println("Summary:")
	log.Println(summary)

	// Save document
	log.Println("Saving document metadata...")

	err = s.documentRepo.Create(document)
	if err != nil {
		log.Println("Document Save Error:", err)
		_ = os.Remove(document.FilePath)
		return "", err
	}

	log.Println("Document saved successfully.")
	log.Println("Document ID:", document.ID)

	// Save summary
	log.Println("Saving summary...")

	summaryModel := &models.Summary{
		DocumentID: document.ID,
		Content:    summary,
		ModelUsed:  "gemini-2.5-flash",
	}

	err = s.summaryRepo.Create(summaryModel)
	if err != nil {
		log.Println("Summary Save Error:", err)
		_ = os.Remove(document.FilePath)
		return "", err
	}

	log.Println("Summary saved successfully.")
	log.Println("Summary ID:", summaryModel.ID)

	log.Println("========== Document Processing Completed ==========")

	return summary, nil
}
