package userservice

import (
	"github.com/sylphritz/go-api-server/pkg/db"
	"github.com/sylphritz/go-api-server/pkg/db/schema"
)

func GetUserById(id uint, user *schema.User) {
	db.DB.First(&user, id)
}

func CreateUser(user *schema.User) error {
	result := db.DB.Create(&user)

	return result.Error
}

func UpdateOrCreateUser(user *schema.User) error {
	result := db.DB.Where(schema.User{Email: user.Email}).Assign(&user).FirstOrCreate(&user)

	db.DB.Updates(user)

	if db.DB.Error != nil {
		return db.DB.Error
	}

	return result.Error
}
