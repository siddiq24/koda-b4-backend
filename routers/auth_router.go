package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/siddiq24/backend-coffee-shop/controllers"
)

func AuthRouter(r *gin.Engine) {
	c := controllers.AuthController{}
	auth := r.Group("auth")
	{
		auth.POST("/register", c.Register)
		auth.POST("/login", c.Login)
	}
}
