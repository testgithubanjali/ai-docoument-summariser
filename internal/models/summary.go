package models

import "gorm.io/gorm"

type Summary struct {
	gorm.Model

	DocumentID uint     `gorm:"not null" json:"document_id"`
	Document   Document `gorm:"foreignKey:DocumentID" json:"document,omitempty"`

	Content   string `gorm:"type:text;not null" json:"content"`
	ModelUsed string `gorm:"size:100;not null" json:"model_used"`
}
