package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/siddiq24/backend-coffee-shop/controllers"
	"github.com/siddiq24/backend-coffee-shop/models"
)

func AuthRouter(r *gin.Engine, pg *pgxpool.Pool) {
	db := models.NewAuth(pg)
	ctrl := controllers.NewAuthController(db)
	auth := r.Group("auth")
	{
		auth.POST("/register", ctrl.Register)
	}
}
