package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/siddiq24/backend-coffee-shop/controllers"
)

func PromoRouter(r *gin.Engine, pg *pgxpool.Pool) {
	p := controllers.Promo{Pg: pg}
	r.GET("/promos", p.GetPromo)
}
