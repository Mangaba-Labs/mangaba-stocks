package config

import (
	"log"

	"github.com/joho/godotenv"
)

// SetupEnvVars set up environment variables
func SetupEnvVars() {
		if err := godotenv.Load(".env"); err != nil {
			log.Fatal("Error loading .env file!")
		}
}