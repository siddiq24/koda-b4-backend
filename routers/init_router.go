package routers

import (
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/siddiq24/backend-coffee-shop/docs"
	"github.com/siddiq24/backend-coffee-shop/middlewares"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	if os.Getenv("ENVIRONMENT") == "production" {
		gin.SetMode(gin.ReleaseMode)
		r.Use(gin.Recovery())
	}
	r.Use(middlewares.Cors())

	InitWelcomeRouter(r)
	AuthRouter(r)
	PromoRouter(r)
	ProductRouter(r)
	OrderRouter(r)
	AdminRouter(r)
	ProfileRouter(r)
	TransactionRouter(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
