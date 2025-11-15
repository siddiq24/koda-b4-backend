package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/siddiq24/backend-coffee-shop/controllers"
)

func InitWelcomeRouter(r *gin.Engine) {
	r.GET("/", controllers.Welcome)
	r.GET("/health", controllers.Health)
}
