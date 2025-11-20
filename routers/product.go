package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/siddiq24/backend-coffee-shop/controllers"
	"github.com/siddiq24/backend-coffee-shop/middlewares"
)

func ProductRouter(r *gin.Engine) {
	c := controllers.ProductsController{}
	r.GET("/products", c.GetAllProducts)
	r.GET("/products/favorite", c.GetAllFavProducts)
	r.GET("/products/:id/recomendation", c.GetRekomendasiById)
	r.GET("/products/:id", c.GetProductByID)

	cart := r.Group("/cart")
	cart.Use(middlewares.AuthMiddleware("user"))
	cart.POST("", c.CreateCart)
	cart.GET("/:id", c.GetProductCart)
	cart.GET("/list", c.GetFullCart)
}
