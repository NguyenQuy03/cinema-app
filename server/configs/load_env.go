package configs

import (
	"log"
	"path/filepath"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load(filepath.Join("../", ".env"))
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}
