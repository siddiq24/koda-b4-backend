package models

import (
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"github.com/siddiq24/backend-coffee-shop/configs"
)

var (
	Pg  *pgxpool.Pool
	Rdb *redis.Client
)

// Helper functions dengan lazy loading
func GetDB() *pgxpool.Pool {
	if Pg == nil {
		// Cek dulu apakah DATABASE_URL ada
		if os.Getenv("DATABASE_URL") == "" {
			log.Println("⚠️ DATABASE_URL not set, skipping database initialization")
			return nil
		}
		Pg = configs.InitPostgres()
	}
	return Pg
}

func GetRedis() *redis.Client {
	if Rdb == nil {
		// Cek dulu apakah REDIS_URL ada
		if os.Getenv("REDIS_URL") == "" && os.Getenv("UPSTASH_REDIS_REST_URL") == "" {
			log.Println("⚠️ Redis URL not set, skipping Redis initialization")
			return nil
		}
		Rdb = configs.Redis()
	}
	return Rdb
}

// Optional: Explicit init function (jangan panggil di init())
func InitConnections() error {
	if os.Getenv("DATABASE_URL") != "" {
		Pg = configs.InitPostgres()
		if Pg == nil {
			return fmt.Errorf("failed to initialize database")
		}
	}

	if os.Getenv("REDIS_URL") != "" || os.Getenv("UPSTASH_REDIS_REST_URL") != "" {
		Rdb = configs.Redis()
	}

	return nil
}
