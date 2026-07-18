package service

import (
	"github.com/testgithubanjali/ai-document-summarizer/internal/models"
	"github.com/testgithubanjali/ai-document-summarizer/internal/repository"
)

type SummaryService struct {
	repo *repository.SummaryRepository
}

func NewSummaryService(repo *repository.SummaryRepository) *SummaryService {
	return &SummaryService{
		repo: repo,
	}
}

func (s *SummaryService) GetAll(userID uint) ([]models.Summary, error) {
	return s.repo.GetAllByUser(userID)
}

func (s *SummaryService) GetByID(id uint) (*models.Summary, error) {
	return s.repo.GetByID(id)
}

func (s *SummaryService) Delete(id uint) error {
	return s.repo.Delete(id)
}
