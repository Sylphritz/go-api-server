package schema

import "gorm.io/gorm"

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
	b.UserID = id
}
