package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
)

var JwtKey []byte

func init() {
	err := godotenv.Load(filepath.Join("..", ".env"))

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	JwtKey = []byte(os.Getenv("JWT_SECRET"))
	if JwtKey == nil {
		log.Fatal("JWT_SECRET env variable not set")
	}
}
