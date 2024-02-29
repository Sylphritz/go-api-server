package schema

import (
	"log"

	"gorm.io/gorm"
)

const UserIdColumnName = "user_id"

type CommonEntity interface {
	GetUserID() uint
	SetUserID(id uint)
}

type BaseEntity struct {
	gorm.Model
	UserID uint `json:"userId"`
}

func (b *BaseEntity) GetUserID() uint {
	return b.UserID
}

func (b *BaseEntity) SetUserID(id uint) {
	log.Printf("before set user id: %v\n", b)
	b.UserID = id
	log.Printf("after set user id: %v\n", b)
}
