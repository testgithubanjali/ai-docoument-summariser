package models

import "gorm.io/gorm"

type Document struct {
	gorm.Model

	UserID   uint   `gorm:"not null"`
	FileName string `gorm:"not null"`
	FilePath string `gorm:"not null"`
	FileType string `gorm:"not null"`
}
