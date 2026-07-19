package repository

import (
	"log"

	"github.com/testgithubanjali/ai-document-summarizer/internal/database"
	"github.com/testgithubanjali/ai-document-summarizer/internal/models"
)

type SummaryRepository struct{}

func NewSummaryRepository() *SummaryRepository {
	return &SummaryRepository{}
}

// Create saves a generated summary
func (r *SummaryRepository) Create(summary *models.Summary) error {

	log.Println("========== Saving Summary ==========")
	log.Printf("DocumentID: %d\n", summary.DocumentID)
	log.Printf("Model Used: %s\n", summary.ModelUsed)
	log.Printf("Summary Length: %d characters\n", len(summary.Content))

	err := database.DB.Create(summary).Error
	if err != nil {
		log.Println("❌ Database Error:", err)
		return err
	}

	log.Printf("✅ Summary saved successfully. Summary ID: %d\n", summary.ID)
	log.Println("===================================")

	return nil
}

// Get all summaries of a user
func (r *SummaryRepository) GetAllByUser(userID uint) ([]models.Summary, error) {

	var summaries []models.Summary

	err := database.DB.
		Joins("JOIN documents ON documents.id = summaries.document_id").
		Where("documents.user_id = ?", userID).
		Preload("Document").
		Find(&summaries).Error

	if err != nil {
		log.Println("❌ Error fetching summaries:", err)
		return nil, err
	}

	log.Printf("✅ Found %d summaries for user %d\n", len(summaries), userID)

	return summaries, nil
}

// Get summary by ID
func (r *SummaryRepository) GetByID(id uint) (*models.Summary, error) {

	var summary models.Summary

	err := database.DB.
		Preload("Document").
		First(&summary, id).Error

	if err != nil {
		log.Println("❌ Error fetching summary:", err)
		return nil, err
	}

	log.Printf("✅ Summary %d fetched successfully\n", summary.ID)

	return &summary, nil
}

// Delete summary
func (r *SummaryRepository) Delete(id uint) error {

	err := database.DB.Delete(&models.Summary{}, id).Error
	if err != nil {
		log.Println("❌ Error deleting summary:", err)
		return err
	}

	log.Printf("✅ Summary %d deleted successfully\n", id)

	return nil
}
