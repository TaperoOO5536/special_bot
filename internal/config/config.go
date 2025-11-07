package config

import (
	// "context"
	// "time"

	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load("pkg/env/.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func GetToken() string {
	return GetEnvDefault("TOKEN", "")
}

func GetEnvDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}