package util

import (
	"golang.org/x/crypto/bcrypt"
)

// GenerateHashedPassword generates an encrypted value of the given password.
func GenerateHashedPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(bytes), err
}

// ValidatePassword checks if the password matches the hashed value.
func ValidatePassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}
