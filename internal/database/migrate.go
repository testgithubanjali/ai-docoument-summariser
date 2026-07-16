package database

import (
	"log"

	"github.com/testgithubanjali/ai-document-summarizer/internal/models"
)

func MigrateDB() {
	err := DB.AutoMigrate(
		&models.User{},
	)

	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Println("✅ Database migration completed")
}
