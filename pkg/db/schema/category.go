package schema

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name        string `gorm:"uniqueIndex;not null" json:"name"`
	Description string `json:"description"`
	Posts       []Post `json:"posts"`
	BlogID      uint   `json:"blogId"`
}
