package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// EnvMongoURI pulls a MongoDB connection string from .env file
func EnvMongoURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("MONGODB_URI")
}

// JWTSecret pulls a secrets string from the .env file
func JWTSecret() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
	return os.Getenv("SECRET_KEY")
}
