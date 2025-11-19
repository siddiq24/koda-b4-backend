package configs

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariable() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("unexepcted env: %x", err)
	}
}
