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
	Pg     *pgxpool.Pool
	PgMsg  string
	pgOnce sync.Once
)

func InitPostgres() *pgxpool.Pool {
	pgOnce.Do(func() {
		if os.Getenv("VERCEL") == "" {
			godotenv.Load()
		}

		connString := os.Getenv("PSQL_URL")

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		config, err := pgxpool.ParseConfig(connString)
		if err != nil {
			PgMsg = "❌ Failed to parse config"
			log.Fatalf("❌ Failed to parse config: %s", err)
		}

		maxConns := 1
		if os.Getenv("ENVIRONMENT") == "development" {
			maxConns = 10
		}

		config.MaxConns = int32(maxConns)
		config.MinConns = 0
		config.MaxConnIdleTime = 30 * time.Second
		config.MaxConnLifetime = 0

		Pg, err = pgxpool.NewWithConfig(context.Background(), config)
		if err != nil {
			PgMsg = "❌ Failed to create pool"
			log.Fatalf("❌ Failed to create pool: %s", err)
		}

		if err := Pg.Ping(ctx); err != nil {
			PgMsg = "❌ Failed to ping database"
			log.Fatalf("❌ Failed to ping database: %s", err)
		}

		PgMsg = "✅ Connected to Postgres successfully"
		log.Println("✅ Connected to Postgres successfully")
	})

	return Pg
}

func GetPostgres() *pgxpool.Pool {
	if Pg == nil {
		return InitPostgres()
	}
	return Pg
}
