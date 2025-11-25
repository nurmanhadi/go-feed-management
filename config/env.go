package config

import (
	"log"

	"github.com/joho/godotenv"
)

func NewEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using OS environment variables")
	}
}
