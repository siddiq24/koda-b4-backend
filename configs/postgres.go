package configs

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitPostgres() *pgxpool.Pool {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_NAME"),
	)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		fmt.Println("✳️ ", err)
	}

	if err := pool.Ping(ctx); err != nil {
		log.Printf("✳️ failed to ping database: %s", err)
		return nil
	}

	log.Println("✳️ Connect to posgres successfully")
	return pool
}
