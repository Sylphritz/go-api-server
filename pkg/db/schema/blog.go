package schema

import (
	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model
	Title       string     `gorm:"uniqueIndex;not null" json:"title"`
	Description string     `json:"description"`
	UserID      uint       `json:"userId"`
	Posts       []Post     `json:"posts"`
	Categories  []Category `json:"categories"`
}
