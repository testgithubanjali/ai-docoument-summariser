package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/testgithubanjali/ai-document-summarizer/internal/config"
)

func GenerateToken(userID uint, email string) (string, error) {

	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(config.GetEnv("JWT_SECRET")))
}
