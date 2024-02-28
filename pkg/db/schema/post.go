package schema

import (
	"time"
)

type Post struct {
	BaseEntity
	Title       string    `gorm:"not null" json:"title"`
	Slug        string    `gorm:"uniqueIndex:idx_post_slug_blog_id;not null" json:"slug"`
	Description string    `gorm:"not null" json:"description"`
	Body        string    `gorm:"not null" json:"body"`
	Published   bool      `json:"published"`
	BlogID      uint      `gorm:"uniqueIndex:idx_post_slug_blog_id;not null" json:"blogId"`
	CategoryID  uint      `json:"categoryId"`
	PublishedAt time.Time `json:"publishedAt"`
}
