package routers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/siddiq24/backend-coffee-shop/docs"
	"github.com/siddiq24/backend-coffee-shop/middlewares"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.Cors())

	AuthRouter(r)
	PromoRouter(r)
	ProductRouter(r)
	OrderRouter(r)
	AdminRouter(r)
	ProfileRouter(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
