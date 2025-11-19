package configs

import (
	"context"
	"fmt"
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
		redisURL := os.Getenv("REDIS_URL")
		if os.Getenv("ENVIRONMENT") == "" {
			godotenv.Load()
			redisURL = os.Getenv("REDIS_URL")
		}

		opt, err := redis.ParseURL(redisURL)
		if err != nil {
			RdbMsg = fmt.Sprintf("❌ Failed to parse Redis URL: %v", err)
			return
		}

		Rdb = redis.NewClient(opt)

		ctx := context.Background()
		if err := Rdb.Ping(ctx).Err(); err != nil {
			RdbMsg = fmt.Sprintf("❌ Failed to connect to Redis: %s", err)
			return
		}

		RdbMsg = "✅ Connected to Redis successfully"
	})

	return Rdb
}

func GetRedis() *redis.Client {
	if Rdb == nil {
		return InitRedis()
	}
	return Rdb
}
