package routers

import (
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/siddiq24/backend-coffee-shop/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter(r *gin.Engine) {
	InitWelcomeRouter(r)
	AuthRouter(r)
	PromoRouter(r)
	ProductRouter(r)
	AdminRouter(r)
	ProfileRouter(r)
	TransactionRouter(r)

	if os.Getenv("ENIRONMENT") == "" {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}
