package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/siddiq24/backend-coffee-shop/models"
)

type Promo struct {
	Pg *pgxpool.Pool
}

func (p *Promo) GetPromo(c *gin.Context) {
	promos, err := models.AllPromo(c)
	if err != nil {
		models.ErrorResponse(c, http.StatusInternalServerError, "Server error", err)
	}

	c.JSON(http.StatusOK, promos)
}
