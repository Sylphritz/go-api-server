package schema

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title       string    `gorm:"not null" json:"title"`
	Slug        string    `gorm:"uniqueIndex;not null" json:"slug"`
	Description string    `gorm:"not null" json:"description"`
	Body        string    `gorm:"not null" json:"body"`
	Published   bool      `json:"published"`
	BlogID      uint      `json:"blogId"`
	CategoryID  uint      `json:"categoryId"`
	UserID      uint      `json:"userId"`
	PublishedAt time.Time `json:"publishedAt"`
}
