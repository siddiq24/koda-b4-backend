package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type ENV struct {
	PORT              string
	PORT_ACCESS       string
	POSTGRES_USER     string
	POSTGRES_PASSWORD string
	POSTGRES_HOST     string
	POSTGRES_PORT     string
	POSTGRES_NAME     string
	JWT_SECRET        string
	REDIS_HOST        string
}

func GetEnv() ENV {
	if os.Getenv("VERCEL") == "" {
		godotenv.Load()
	}
	return ENV{
		PORT:              os.Getenv("PORT"),
		PORT_ACCESS:       os.Getenv("PORT_ACCESS"),
		POSTGRES_USER:     os.Getenv("POSTGRES_USER"),
		POSTGRES_PASSWORD: os.Getenv("POSTGRES_PASSWORD"),
		POSTGRES_HOST:     os.Getenv("POSTGRES_HOST"),
		POSTGRES_PORT:     os.Getenv("POSTGRES_PORT"),
		POSTGRES_NAME:     os.Getenv("POSTGRES_NAME"),
		JWT_SECRET:        os.Getenv("JWT_SECRET"),
		REDIS_HOST:        os.Getenv("REDIS_HOST"),
	}
}
