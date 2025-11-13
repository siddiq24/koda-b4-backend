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
		if os.Getenv("VERCEL") == "" {
			godotenv.Load()
		}

		redisURL := os.Getenv("REDIS_URL")
		if redisURL == "" {
			log.Fatal("❌ REDIS_URL not found in environment variables")
		}

		// ✅ Gunakan redis.ParseURL agar format rediss:// bisa dibaca
		opt, err := redis.ParseURL(redisURL)
		if err != nil {
			log.Fatalf("❌ Failed to parse Redis URL: %v", err)
		}

		redisClient = redis.NewClient(opt)

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
