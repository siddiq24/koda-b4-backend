package configs

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func Redis() *redis.Client {
	godotenv.Load()
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: "",
		DB:       0,
	})
	return client
}
