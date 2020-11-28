package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvVar(key string) string {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	return os.Getenv(key)
}
