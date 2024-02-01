package schema

import (
	"github.com/sylphritz/go-api-server/pkg/constants"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email                   string         `gorm:"uniqueIndex;not null" json:"email"`
	Password                string         `json:"-"`
	GoogleOAuthToken        string         `json:"-"`
	GoogleOAuthRefreshToken string         `json:"-"`
	Role                    constants.Role `json:"role"`
	CustomerID              string         `json:"-"`
	SubscriptionID          string         `json:"-"`
	PaymentMethodID         string         `json:"-"`
	Blogs                   []Blog         `json:"blogs"`
	Posts                   []Post         `json:"posts"`
}
