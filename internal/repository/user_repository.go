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
func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User

	err := database.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
