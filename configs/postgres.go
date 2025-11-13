package configs

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var (
	pgPool *pgxpool.Pool
	pgOnce sync.Once
)

func InitPostgres() *pgxpool.Pool {
	pgOnce.Do(func() {
		if os.Getenv("VERCEL") == "" {
			godotenv.Load()
		}

		// Support both individual env vars and DATABASE_URL
		connString := os.Getenv("PSQL_URL")

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Parse config untuk custom settings
		config, err := pgxpool.ParseConfig(connString)
		if err != nil {
			log.Fatalf("❌ Failed to parse config: %s", err)
		}

		// Optimized untuk serverless (Vercel)
		// Untuk local development, bisa diubah via env var
		maxConns := 1
		if os.Getenv("ENVIRONMENT") == "development" {
			maxConns = 10
		}

		config.MaxConns = int32(maxConns)
		config.MinConns = 0
		config.MaxConnIdleTime = 30 * time.Second
		config.MaxConnLifetime = 0

		pgPool, err = pgxpool.NewWithConfig(context.Background(), config)
		if err != nil {
			log.Fatalf("❌ Failed to create pool: %s", err)
		}

		if err := pgPool.Ping(ctx); err != nil {
			log.Fatalf("❌ Failed to ping database: %s", err)
		}

		log.Println("✅ Connected to Postgres successfully")
	})

	return pgPool
}

// GetPool returns the singleton pool instance
func GetPool() *pgxpool.Pool {
	if pgPool == nil {
		return InitPostgres()
	}
	return pgPool
}
