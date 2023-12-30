package schema

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email                   string `gorm:"uniqueIndex;not null"`
	Password                string `json:"-"`
	GoogleOAuthToken        string `json:"-"`
	GoogleOAuthRefreshToken string `json:"-"`
}
