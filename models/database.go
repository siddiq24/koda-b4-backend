package models

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"github.com/siddiq24/backend-coffee-shop/configs"
)

var (
	Pg  *pgxpool.Pool
	Rdb *redis.Client
)

func init() {
	Pg = configs.InitPostgres()
	Rdb = configs.Redis()
}

// Optional: Helper functions
func GetDB() *pgxpool.Pool {
	if Pg == nil {
		Pg = configs.InitPostgres()
	}
	return Pg
}

func GetRedis() *redis.Client {
	if Rdb == nil {
		Rdb = configs.Redis()
	}
	return Rdb
}
