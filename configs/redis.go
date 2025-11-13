package configs

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var (
	redisClient *redis.Client
	redisOnce   sync.Once
)

func Redis() *redis.Client {
	redisOnce.Do(func() {
		godotenv.Load()

		// Support Upstash Redis REST API atau standard Redis
		redisURL := os.Getenv("REDIS_URL")
		redisPassword := os.Getenv("REDIS_PASSWORD")

		redisClient = redis.NewClient(&redis.Options{
			Addr:     redisURL,
			Password: redisPassword,
			DB:       0,
		})

		ctx := context.Background()
		if err := redisClient.Ping(ctx).Err(); err != nil {
			log.Fatalf("❌ Failed to connect to Redis: %s", err)
		}

		log.Println("✅ Connected to Redis successfully")
	})

	return redisClient
}

// GetRedis returns the singleton redis client
func GetRedis() *redis.Client {
	if redisClient == nil {
		return Redis()
	}
	return redisClient
}
