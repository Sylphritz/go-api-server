package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Println(".env not found. Using default values.")
		return
	}

	log.Println(".env loaded")
}
