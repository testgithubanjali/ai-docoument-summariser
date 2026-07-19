package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/testgithubanjali/ai-document-summarizer/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	log.Println("===== Database Configuration =====")
	log.Println("DB_HOST:", config.GetEnv("DB_HOST"))
	log.Println("DB_PORT:", config.GetEnv("DB_PORT"))
	log.Println("DB_USER:", config.GetEnv("DB_USER"))
	log.Println("DB_NAME:", config.GetEnv("DB_NAME"))

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.GetEnv("DB_HOST"),
		config.GetEnv("DB_USER"),
		config.GetEnv("DB_PASSWORD"),
		config.GetEnv("DB_NAME"),
		config.GetEnv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
	}

	DB = db

	// Verify which database is actually connected
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("❌ Failed to get SQL DB:", err)
	}

	var currentDatabase string
	err = sqlDB.QueryRow("SELECT current_database()").Scan(&currentDatabase)
	if err != nil && err != sql.ErrNoRows {
		log.Fatal("❌ Failed to get current database:", err)
	}

	log.Println("✅ Connected Database:", currentDatabase)
	log.Println("✅ Database connected successfully")
}
