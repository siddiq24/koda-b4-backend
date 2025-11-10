package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/siddiq24/backend-coffee-shop/models"
)

type AdminController struct {
	Admin models.Admin
}

// GetAllProducts godoc
// @Summary     Get all products
// @Description Mendapatkan daftar produk dengan paginasi, limit, dan filter berdasarkan nama & kategori
// @Tags        D_admin
// @Accept      json
// @Accept      x-www-form-urlencoded
// @Produce     json
// @Param       Page        	query 	int    	false 	"Nomor halaman (default: 1)" default(1) minimum(1)
// @Param       Search      	query 	string 	false 	"Filter berdasarkan nama atau deskripsi produk"
// @Security	BearerAuth
// @Success     200  {object}  models.JSON_Response{result=[]object} "Berhasil mendapatkan daftar produk"
// @Failure     400  {object}  models.JSON_Response "Invalid query parameters"
// @Failure     500  {object}  models.JSON_Response "Internal server error"
// @Router      /admin/products [get]
func (a AdminController) GetAllProducts(c *gin.Context) {
	var req models.Param
	if err := c.ShouldBindQuery(&req); err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, "Terjadi kesalahan saat Binding query param", err.Error())
		return
	}
	products, err := a.Admin.GetAllProducts(c, req.Page, req.Search)
	if err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, "gagal mendapatkan products", err.Error())
		return
	}
	c.JSON(http.StatusOK, models.JSON_Response{
		Success: true,
		Message: "berhasil mendapatkan products",
		Result:  products,
	})
}
