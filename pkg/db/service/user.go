package service

import (
	"github.com/sylphritz/go-api-server/pkg/db"
	"github.com/sylphritz/go-api-server/pkg/db/schema"
)

func CreateUser(user *schema.User) error {
	result := db.DB.Create(&user)

	return result.Error
}
