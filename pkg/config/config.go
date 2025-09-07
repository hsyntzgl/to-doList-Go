package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Config(key string) string {

	err := godotenv.Load("../../.env")

	if err != nil {
		log.Fatalf("Error load .env file: %v", err)
	}

	value := os.Getenv(key)

	if value == "" {
		panic("the key doesn't exist in .env file")
	}
	return value
}
