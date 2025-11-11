package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/siddiq24/backend-coffee-shop/models"
)

type ProductsController struct {
	Product *models.Product
}

// GetAllProducts godoc
// @Summary      Get all products
// @Description  Mendapatkan daftar produk dengan paginasi, limit, dan filter berdasarkan nama & kategori
// @Tags         products
// @Accept       json
// @Accept       x-www-form-urlencoded
// @Produce      json
// @Param        Page        	query 	int    	false 	"Nomor halaman (default: 1)" default(1) minimum(1)
// @Param        Search      	query 	string 	false 	"Filter berdasarkan nama atau deskripsi produk"
// @Param        CategoryId 	query 	int    	false 	"Filter berdasarkan ID kategori" minimum(1)
// @Param        MinPrice   	query 	int    	false 	"Filter harga minimum" minimum(0)
// @Param        MaxPrice   	query 	int    	false 	"Filter harga maksimum" minimum(0)
// @Success      200  {object}  models.JSON_Response{result=[]object} "Berhasil mendapatkan daftar produk"
// @Failure      400  {object}  models.JSON_Response "Invalid query parameters"
// @Failure      500  {object}  models.JSON_Response "Internal server error"
// @Router       /products [get]
func (pc *ProductsController) GetAllProducts(c *gin.Context) {
	var prm models.Product_Params

	if err := c.ShouldBindQuery(&prm); err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, "invalid query params", err.Error())
		return
	}

	value, err := models.Rdb.Get(c.Request.Context(), "users").Result()
	var results []models.Product_ress

	// jika redis kosong, ambil dari postgres
	if value == "" || err != nil {
		products, err := pc.Product.AllProductFiltered(c, prm)
		if err != nil {
			models.ErrorResponse(c, http.StatusInternalServerError, "failed to get products", err.Error())
			return
		}
		prod, _ := json.Marshal(products)

		if err := models.Rdb.Set(c.Request.Context(), "users", prod, (time.Duration(15) * time.Second)).Err(); err != nil {
			models.ErrorResponse(c, http.StatusBadRequest, "Gagal Set products ke redis", err.Error())
			return
		}
		results = append(results, products...)
	} else {
		if err := json.Unmarshal([]byte(value), &results); err != nil {
			models.ErrorResponse(c, http.StatusBadRequest, "Gagal  unmarshal", err.Error())
			return
		}
	}

	c.JSON(200, models.JSON_Response{
		Success: true,
		Message: "success get products",
		Result:  results,
	})
}
