package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/siddiq24/backend-coffee-shop/controllers"
)

func PromoRouter(r *gin.Engine) {
	c := controllers.Promo{}
	r.GET("/promos", c.GetPromo)
}
