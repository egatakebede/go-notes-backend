package models

import (
	"time"

	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
