package configs

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var (
	Pg     *pgxpool.Pool
	PgMsg  string
	pgOnce sync.Once
)

// InitDB inisialisasi koneksi database
func InitDB() *pgxpool.Pool {
	pgOnce.Do(func() {
		databaseURL := os.Getenv("DATABASE_URL")
		if os.Getenv("ENVIRONMENT") == "" {
			godotenv.Load()
			databaseURL = os.Getenv("DATABASE_URL")
		}

		config, err := pgxpool.ParseConfig(databaseURL)
		if err != nil {
			PgMsg = fmt.Sprintf("❌ Unable to parse DATABASE_URL: %v\n", err)
			return
		}

		// Konfigurasi pool connection
		config.MaxConns = 10
		config.MinConns = 2

		pool, err := pgxpool.NewWithConfig(context.Background(), config)
		if err != nil {
			PgMsg = fmt.Sprintf("❌ Unable to create connection pool: %v\n", err)
			return
		}

		// Test koneksi
		if err := pool.Ping(context.Background()); err != nil {
			PgMsg = fmt.Sprintf("❌ Unable to ping database: %v\n", err)
			return
		}

		Pg = pool
		PgMsg = "✅ Database connected successfully"
	})

	return Pg
}

func GetDB() *pgxpool.Pool {
	if Pg == nil {
		return InitDB()
	}
	return Pg
}
