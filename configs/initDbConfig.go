package configs

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitDbConfigs() *pgxpool.Pool {
	pg := InitPostgres()

	return pg
}
