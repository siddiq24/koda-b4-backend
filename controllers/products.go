package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/siddiq24/backend-coffee-shop/libs"
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
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	min, _ := strconv.Atoi(c.Query("minPrice"))
	max, _ := strconv.Atoi(c.Query("maxPrice"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	_, ok := c.GetQuery("asc")
	prm := models.Product_Params{
		Page:       page,
		Search:     c.Query("search"),
		CategoryId: c.QueryArray("cat"),
		MinPrice:   uint64(min),
		MaxPrice:   uint64(max),
		ShortBy:    c.Query("shortBy"),
		Asc:        ok,
		Limit:      limit,
	}

	var value string
	var err error
	var results []models.Product_ress

	if len(c.Request.URL.Query()) == 0 {
		value, err = models.Rdb.Get(c.Request.Context(), "users").Result()
		fmt.Println("get redis")
	}

	// jika redis kosong, ambil dari postgres
	if value == "" || err != nil {
		products, err := pc.Product.AllProductFiltered(c, prm)
		if err != nil {
			models.ErrorResponse(c, http.StatusInternalServerError, "failed to get products", err.Error())
			return
		}
		if len(c.Request.URL.Query()) == 0 {
			prod, _ := json.Marshal(products)
			fmt.Println("set redis")

			if err := models.Rdb.Set(c.Request.Context(), "users", prod, (time.Duration(15) * time.Second)).Err(); err != nil {
				models.ErrorResponse(c, http.StatusBadRequest, "Gagal Set products ke redis", err.Error())
				return
			}
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

func (pc *ProductsController) GetAllFavProducts(c *gin.Context) {
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, "Invalid Query", err.Error())
		return
	}

	results, err := pc.Product.FavProducts(c, limit)
	if err != nil {
		models.ErrorResponse(c, http.StatusInternalServerError, "Error from get Fav products ", err.Error())
		return
	}

	c.JSON(200, models.JSON_Response{
		Success: true,
		Message: "success get products",
		Result:  results,
	})
}

func (pc *ProductsController) GetRekomendasiById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		models.ErrorResponse(c, http.StatusInternalServerError, "Invalid Param ", err.Error())
		return
	}
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		models.ErrorResponse(c, http.StatusInternalServerError, "Invalid Param ", err.Error())
		return
	}

	ress, err := pc.Product.GetRecommendation(c.Request.Context(), id, limit)
	if err != nil {
		models.ErrorResponse(c, http.StatusInternalServerError, "Invalid Param ", err.Error())
		return
	}

	c.JSON(http.StatusOK, models.JSON_Response{
		Success: true,
		Message: "Get recomendation Products successfully",
		Result:  ress,
	})
}

func (pc *ProductsController) CreateCart(c *gin.Context) {
	claim, err := libs.VerifyJwt((c.Request.Header.Get("Authorization")[7:]))
	if err != nil {
		models.ErrorResponse(c, http.StatusNetworkAuthenticationRequired, "Unauthorize", err.Error())
	}
	id := int((*claim)["id"].(float64))
	var req models.Cart_Request
	req.UserId = id
	if err := c.ShouldBind(&req); err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, "Request Invalid", err.Error())
		return
	}
	ress, err := pc.Product.CreateCart(c.Request.Context(), req)
	if err != nil {
		models.ErrorResponse(c, http.StatusInternalServerError, "Internal service error", err.Error())
		return
	}

	c.JSON(http.StatusOK, models.JSON_Response{
		Success: true,
		Message: "Add product to cart successfully",
		Result:  ress,
	})
}

func (pc *ProductsController) GetProductCart(c *gin.Context) {
	claim, err := libs.VerifyJwt((c.Request.Header.Get("Authorization")[7:]))
	if err != nil {
		models.ErrorResponse(c, http.StatusNetworkAuthenticationRequired, "Unauthorize", err.Error())
	}
	id := int((*claim)["id"].(float64))
	ress, err := pc.Product.GetProductCart(c.Request.Context(), id)
	if err != nil {
		models.ErrorResponse(c, http.StatusInternalServerError, "Internal service error", err.Error())
		return
	}

	c.JSON(http.StatusOK, models.JSON_Response{
		Success: true,
		Message: "Get product from cart successfully",
		Result:  ress,
	})
}

func (pc *ProductsController) GetfullCard(c *gin.Context) {
	claim, err := libs.VerifyJwt((c.Request.Header.Get("Authorization")[7:]))
	if err != nil {
		models.ErrorResponse(c, http.StatusNetworkAuthenticationRequired, "Unauthorize", err.Error())
		return
	}
	id := int((*claim)["id"].(float64))

	ress, err := pc.Product.GetListCart(c.Request.Context(), id)
	if err != nil {
		models.ErrorResponse(c, http.StatusInternalServerError, "Server error", err.Error())
		return
	}

	c.JSON(http.StatusOK, models.JSON_Response{
		Success: true,
		Message: "Get all list cart successfully",
		Result:  ress,
	})

}
