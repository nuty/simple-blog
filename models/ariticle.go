package models

import (
	"time"
)

type Article struct {
	ID        uint        `gorm:"primaryKey" json:"id"`
  	CreatedAt time.Time   `json:"created_at"`
	Slug        string    `gorm:"uniqueIndex;not null" json:"slug"`
	Title       string    `gorm:"not null" json:"title"`
	Description string    `json:"description"`
	Content     string    `json:"content"`
}


