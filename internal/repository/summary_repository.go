package repository

import (
	"github.com/testgithubanjali/ai-document-summarizer/internal/database"
	"github.com/testgithubanjali/ai-document-summarizer/internal/models"
)

type SummaryRepository struct{}

func NewSummaryRepository() *SummaryRepository {
	return &SummaryRepository{}
}

func (r *SummaryRepository) Create(summary *models.Summary) error {
	return database.DB.Create(summary).Error
}

func (r *SummaryRepository) GetAllByUser(userID uint) ([]models.Summary, error) {
	var summaries []models.Summary

	err := database.DB.
		Joins("JOIN documents ON documents.id = summaries.document_id").
		Where("documents.user_id = ?", userID).
		Preload("Document").
		Find(&summaries).Error

	return summaries, err
}

func (r *SummaryRepository) GetByID(id uint) (*models.Summary, error) {
	var summary models.Summary

	err := database.DB.
		Preload("Document").
		First(&summary, id).Error

	if err != nil {
		return nil, err
	}

	return &summary, nil
}

func (r *SummaryRepository) Delete(id uint) error {
	return database.DB.Delete(&models.Summary{}, id).Error
}
