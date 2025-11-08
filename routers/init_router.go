package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/siddiq24/backend-coffee-shop/docs"
	"github.com/siddiq24/backend-coffee-shop/middlewares"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter(pg *pgxpool.Pool) *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.Cors())

	AuthRouter(r, pg)
	PromoRouter(r, pg)
	ProductRouter(r, pg)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
