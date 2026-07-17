package service

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/testgithubanjali/ai-document-summarizer/internal/models"
	"github.com/testgithubanjali/ai-document-summarizer/internal/repository"
)

var ErrEmailAlreadyExists = errors.New("email already exists")

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) Register(user *models.User) error {

	_, err := s.userRepo.FindByEmail(user.Email)

	if err == nil {
		return ErrEmailAlreadyExists
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	return s.userRepo.Create(user)
}
