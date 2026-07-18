package models

import "gorm.io/gorm"

type Summary struct {
	gorm.Model

	DocumentID uint   `gorm:"not null"`
	Content    string `gorm:"type:text;not null"`
	ModelUsed  string `gorm:"not null"`
}
