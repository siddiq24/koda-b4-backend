package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/siddiq24/backend-coffee-shop/controllers"
	"github.com/siddiq24/backend-coffee-shop/middlewares"
)

func ProfileRouter(r *gin.Engine) {
	prof := r.Group("/profile")
	prof.Use(middlewares.AuthMiddleware("user"))
	prof.PATCH("", controllers.UpdateProfile)
	prof.PATCH("/image", controllers.UpdateProfileImage)
	prof.GET("", controllers.GetProfileInfo)
}
