package service

import (
	"github.com/sylphritz/go-api-server/pkg/db"
	"github.com/sylphritz/go-api-server/pkg/db/schema"
)

func CreateUser(user *schema.User) error {
	result := db.DB.Create(&user)

	return result.Error
}

func UpdateOrCreateUser(user *schema.User) error {
	var foundUser schema.User
	result := db.DB.Where(schema.User{Email: user.Email}).Assign(&user).FirstOrCreate(&foundUser)

	db.DB.Updates(foundUser)

	if db.DB.Error != nil {
		return db.DB.Error
	}

	return result.Error
}
