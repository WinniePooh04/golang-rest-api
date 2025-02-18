package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	envFile := ".env"
	if os.Getenv("GO_ENV") == "test" {
		envFile = ".env.test"
	}

	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
