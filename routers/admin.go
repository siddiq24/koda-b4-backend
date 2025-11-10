package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/siddiq24/backend-coffee-shop/controllers"
	"github.com/siddiq24/backend-coffee-shop/middlewares"
)

func AdminRouter(r *gin.Engine) {
	c := controllers.AdminController{}
	admin := r.Group("/admin")
	admin.Use(middlewares.AuthMiddleware("admin"))
	{
		admin.GET("products", c.GetAllProducts)
	}
}
