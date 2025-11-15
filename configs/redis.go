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
	Rdb       *redis.Client
	RdbMsg    string
	redisOnce sync.Once
)

func InitRedis() *redis.Client {
	redisOnce.Do(func() {
		if os.Getenv("VERCEL") == "" {
			godotenv.Load()
		}

		redisURL := os.Getenv("REDIS_URL")
		if redisURL == "" {
			RdbMsg = "❌ REDIS_URL not found in environment variables"
			log.Fatal("❌ REDIS_URL not found in environment variables")
		}

		opt, err := redis.ParseURL(redisURL)
		if err != nil {
			RdbMsg = "❌ Failed to parse Redis URL"
			log.Fatalf("❌ Failed to parse Redis URL: %v", err)
		}

		Rdb = redis.NewClient(opt)

		ctx := context.Background()
		if err := Rdb.Ping(ctx).Err(); err != nil {
			RdbMsg = "❌ Failed to connect to Redis"
			log.Fatalf("❌ Failed to connect to Redis: %s", err)
		}

		RdbMsg = "✅ Connected to Redis successfully"
		log.Println("✅ Connected to Redis successfully")
	})

	return Rdb
}

func GetRedis() *redis.Client {
	if Rdb == nil {
		return InitRedis()
	}
	return Rdb
}
