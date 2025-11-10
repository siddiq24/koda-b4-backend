package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/siddiq24/backend-coffee-shop/controllers"
	"github.com/siddiq24/backend-coffee-shop/middlewares"
)

func OrderRouter(r *gin.Engine) {
	order := controllers.OrderController{}

	r_order := r.Group("/order").Use(middlewares.AuthMiddleware("user"))

	r_order.POST("", order.CreateOrder)
}
