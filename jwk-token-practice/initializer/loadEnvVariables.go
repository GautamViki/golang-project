package initializer

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	if err := godotenv.Load(); err != nil {
log.Fatal("Error loading Env file.")
	}
}
