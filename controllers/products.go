package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/siddiq24/backend-coffee-shop/models"
)

type ProductsController struct {
	Product *models.Product
}

func NewProductController(pg *pgxpool.Pool) *ProductsController {
	return &ProductsController{
		Product: &models.Product{Pg: pg},
	}
}

func (pc *ProductsController) GetAllProducts(c *gin.Context) {
	var prm models.Product_Params

	if err := c.ShouldBindQuery(&prm); err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "invalid query params",
			"error":   err.Error(),
		})
		return
	}

	products, err := pc.Product.AllProductFiltered(c, prm)
	if err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": "failed to get products",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"data":    products,
	})
}
