package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/siddiq24/backend-coffee-shop/controllers"
	"github.com/siddiq24/backend-coffee-shop/middlewares"
)

func TransactionRouter(r *gin.Engine) {
	var c controllers.TransactionControoller
	transactions := r.Group("/transactions")
	transactions.Use(middlewares.AuthMiddleware("user"))
	transactions.POST("", c.CreateTransactions)
	transactions.GET("/history", c.GetHistory)
}
