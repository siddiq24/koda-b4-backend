package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/siddiq24/backend-coffee-shop/controllers"
	"github.com/siddiq24/backend-coffee-shop/middlewares"
)

func OrderRouter(r *gin.Engine, pg *pgxpool.Pool) {
	order := controllers.NewOrderController(pg)

	r_order := r.Group("/order").Use(middlewares.AuthMiddleware("user"))

	r_order.POST("", order.CreateOrder)
}
