package repository

import (
	"github.com/testgithubanjali/ai-document-summarizer/internal/database"
	"github.com/testgithubanjali/ai-document-summarizer/internal/models"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) Create(user *models.User) error {
	return database.DB.Create(user).Error
}
