package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/siddiq24/backend-coffee-shop/controllers"
)

func ProductRouter(r *gin.Engine) {
	c := controllers.ProductsController{}
	r.GET("/products", c.GetAllProducts)
}
