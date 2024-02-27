package schema

import (
	"time"
)

type Post struct {
	BaseEntity
	Title       string    `gorm:"not null" json:"title"`
	Slug        string    `gorm:"uniqueIndex;not null" json:"slug"`
	Description string    `gorm:"not null" json:"description"`
	Body        string    `gorm:"not null" json:"body"`
	Published   bool      `json:"published"`
	BlogID      uint      `json:"blogId"`
	CategoryID  uint      `json:"categoryId"`
	PublishedAt time.Time `json:"publishedAt"`
}
