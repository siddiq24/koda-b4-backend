package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/siddiq24/backend-coffee-shop/controllers"
	"github.com/siddiq24/backend-coffee-shop/models"
)

func ProductRouter(r *gin.Engine, pg *pgxpool.Pool) {
	p := controllers.ProductsController{Product: &models.Product{Pg: pg}}
	r.GET("/products", p.GetAllProducts)
}
