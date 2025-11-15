package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVariable() {
	if os.Getenv("VERCEL") == "" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("unexepcted env: %x", err)
		}
	}
}
