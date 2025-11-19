package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/siddiq24/backend-coffee-shop/controllers"
	"github.com/siddiq24/backend-coffee-shop/middlewares"
)

func AuthRouter(r *gin.Engine) {
	c := controllers.AuthController{}
	auth := r.Group("auth")
	{
		auth.POST("/register", c.Register)
		auth.POST("/login", c.Login)
		auth.POST("/forgot-password", c.ForgotPassword)
		auth.POST("/insert-pin", c.ValidatePin)
		auth.POST("/set-new-password", c.SetNewPassword)
		auth.POST("/logout", c.Logout).Use(middlewares.AuthMiddleware("all"))
	}
}
