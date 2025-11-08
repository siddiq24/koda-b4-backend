package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/siddiq24/backend-coffee-shop/middlewares"
)

func InitRouter(pg *pgxpool.Pool) *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.Cors())

	AuthRouter(r, pg)

	return r
}
