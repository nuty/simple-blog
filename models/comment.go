package models

import (
	"time"
)

type Comment struct {
	ID               uint      `gorm:"primaryKey" json:"id"`
	CreatedAt        time.Time `json:"created_at"`
	ArticleID        uint      `json:"article_id"`
	ParentCommentID  *uint     `gorm:"index;foreignKey:ID;constraint:OnDelete:CASCADE;" json:"parent_comment_id"`
	Username         string    `gorm:"not null;size:100" json:"username"`
	Email            string    `gorm:"not null;size:200" json:"email"`
	Content          string    `gorm:"type:text;not null" json:"content"`
	Children         []Comment `gorm:"foreignKey:ParentCommentID;references:ID;constraint:OnDelete:CASCADE;" json:"replies"`
}
