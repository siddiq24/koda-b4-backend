package controllers

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/siddiq24/backend-coffee-shop/configs"
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
		value, err = configs.GetRedis().Get(c, "products:page1").Result()
		fmt.Println("get redis")
	}
	totalPages := 0

	// jika redis kosong, ambil dari postgres
	if value == "" || err != nil {
		products := []models.Product_ress{}
		fmt.Println("redis kosong")
		products, totalPages, err = pc.Product.AllProductFiltered(c, prm)
		if err != nil {
			models.ErrorResponse(c, http.StatusInternalServerError, "failed to get products", err.Error())
			return
		}
		if len(c.Request.URL.Query()) == 0 {
			prod, _ := json.Marshal(products)
			fmt.Println("set redis")

			if err := configs.GetRedis().Set(c, "products:page1", prod, (time.Duration(15) * time.Second)).Err(); err != nil {
				models.ErrorResponse(c, http.StatusBadRequest, "Gagal Set products ke redis", err.Error())
				return
			}
		}
		results = products
	} else {
		if err := json.Unmarshal([]byte(value), &results); err != nil {
			models.ErrorResponse(c, http.StatusBadRequest, "Gagal unmarshal", err.Error())
			return
		}
	}

	fmt.Println(c.Request.URL.String())
	prev := libs.BuildPageURL(c, 2)
	next := ""

	if page > 1 {
		prev = libs.BuildPageURL(c, page-1)
	}

	if page < totalPages {
		next = libs.BuildPageURL(c, page+1)
	}
	fmt.Println(prev)
	fmt.Println(next)

	c.JSON(200, models.ProductsRessponse{
		Success:    true,
		Message:    "success get products",
		Page:       page,
		NextPage:   next,
		PrevPage:   prev,
		TotalPages: totalPages,
		Result:     results,
	})
}

// GetProductByID godoc
// @Summary      Get product by ID
// @Description  Mendapatkan detail produk berdasarkan ID
// @Tags         products
// @Accept       json
// @Accept       x-www-form-urlencoded
// @Produce      json
// @Param        id   path      int  true  "Product ID"
// @Success      200  {object}  models.JSON_Response{result=[]models.Product_ress} "Berhasil mendapatkan detail produk"
// @Failure      400  {object}  models.JSON_Response "Invalid product ID"
// @Failure      404  {object}  models.JSON_Response "Product not found"
// @Failure      500  {object}  models.JSON_Response "Internal server error"
// @Router       /products/{id} [get]
func (pc *ProductsController) GetProductByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, "Invalid product ID", err.Error())
		return
	}

	// PERBAIKAN: Gunakan GetProductByID yang mengembalikan single product
	product, err := pc.Product.GetProductByID(c, int64(id))
	if err != nil {
		if err.Error() == "product not found" {
			models.ErrorResponse(c, http.StatusNotFound, "Product not found", err.Error())
			return
		}
		models.ErrorResponse(c, http.StatusInternalServerError, "Failed to get product", err.Error())
		return
	}

	// PERBAIKAN: Return sebagai array untuk kompatibilitas dengan frontend
	c.JSON(200, models.JSON_Response{
		Success: true,
		Message: "success get product",
		Result:  []models.Product_ress{product}, // Wrap dalam array
	})
}

// GetAllFavProducts godoc
// @Summary      Get favorite products
// @Description  Mendapatkan daftar produk favorit
// @Tags         products
// @Accept       json
// @Accept       x-www-form-urlencoded
// @Produce      json
// @Param        limit   query     int  false  "Limit jumlah produk (default: 10)"
// @Success      200  {object}  models.JSON_Response{result=[]models.Product_ress} "Berhasil mendapatkan produk favorit"
// @Failure      500  {object}  models.JSON_Response "Internal server error"
// @Router       /products/favorites [get]
func (pc *ProductsController) GetAllFavProducts(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if limit <= 0 {
		limit = 10
	}

	results, err := pc.Product.FavProducts(c, limit)
	if err != nil {
		models.ErrorResponse(c, http.StatusInternalServerError, "Error getting favorite products", err.Error())
		return
	}

	c.JSON(200, models.JSON_Response{
		Success: true,
		Message: "success get favorite products",
		Result:  results,
	})
}

// GetRekomendasiById godoc
// @Summary      Get product recommendations
// @Description  Mendapatkan rekomendasi produk berdasarkan product ID
// @Tags         products
// @Accept       json
// @Accept       x-www-form-urlencoded
// @Produce      json
// @Param        id     path      int  true  "Product ID"
// @Param        limit  query     int  false "Limit jumlah rekomendasi (default: 10)"
// @Success      200  {object}  models.JSON_Response{result=[]models.Product_ress} "Berhasil mendapatkan rekomendasi produk"
// @Failure      400  {object}  models.JSON_Response "Invalid parameters"
// @Failure      500  {object}  models.JSON_Response "Internal server error"
// @Router       /products/{id}/recommendations [get]
func (pc *ProductsController) GetRekomendasiById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, "Invalid product ID", err.Error())
		return
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "4"))
	if err != nil {
		limit = 10
	}

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		limit = 10
	}

	if limit <= 0 {
		limit = 10
	}

	results, totalCount, err := pc.Product.GetRecommendation(c.Request.Context(), id, page, limit)
	if err != nil {
		models.ErrorResponse(c, http.StatusInternalServerError, "Failed to get recommendations", err.Error())
		return
	}

	totalPages := int(math.Ceil(float64(totalCount) / float64(limit)))

	c.JSON(http.StatusOK, models.JSON_Response{
		Success:   true,
		Message:   "Get product recommendations successfully",
		TotalPage: totalPages,
		Result:    results,
	})
}

// CreateCart godoc
// @Summary      Add product to cart
// @Description  Menambahkan produk ke keranjang belanja
// @Tags         cart
// @Accept       json
// @Accept       x-www-form-urlencoded
// @Produce      json
// @Security     ApiKeyAuth
// @Param        request  body      models.CartRequest  true  "Cart request"
// @Success      200  {object}  models.JSON_Response{result=models.CartItem} "Berhasil menambahkan ke keranjang"
// @Failure      400  {object}  models.JSON_Response "Invalid request"
// @Failure      401  {object}  models.JSON_Response "Unauthorized"
// @Failure      500  {object}  models.JSON_Response "Internal server error"
// @Router       /cart [post]
func (pc *ProductsController) CreateCart(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	if len(authHeader) < 8 {
		models.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized", "Invalid authorization header")
		return
	}

	claim, err := libs.VerifyJwt(authHeader[7:])
	if err != nil {
		models.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized", err.Error())
		return
	}

	id := int((*claim)["id"].(float64))
	var req models.CartRequest
	req.UserId = id

	if err := c.ShouldBind(&req); err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, "Invalid request", err.Error())
		return
	}

	result, err := pc.Product.AddToCart(c.Request.Context(), req)
	if err != nil {
		models.ErrorResponse(c, http.StatusInternalServerError, "Internal service error", err.Error())
		return
	}

	c.JSON(http.StatusOK, models.JSON_Response{
		Success: true,
		Message: "Add product to cart successfully",
		Result:  result,
	})
}

// GetProductCart godoc
// @Summary      Get cart item by ID
// @Description  Mendapatkan item keranjang berdasarkan ID
// @Tags         cart
// @Accept       json
// @Accept       x-www-form-urlencoded
// @Produce      json
// @Security     ApiKeyAuth
// @Param        id   path      int  true  "Cart Item ID"
// @Success      200  {object}  models.JSON_Response{result=models.CartItem} "Berhasil mendapatkan item keranjang"
// @Failure      400  {object}  models.JSON_Response "Invalid cart ID"
// @Failure      401  {object}  models.JSON_Response "Unauthorized"
// @Failure      500  {object}  models.JSON_Response "Internal server error"
// @Router       /cart/{id} [get]
func (pc *ProductsController) GetProductCart(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	if len(authHeader) < 8 {
		models.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized", "Invalid authorization header")
		return
	}

	claim, err := libs.VerifyJwt(authHeader[7:])
	if err != nil {
		models.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized", err.Error())
		return
	}

	uid := int((*claim)["id"].(float64))
	pid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, "Invalid cart ID", err.Error())
		return
	}

	result, err := pc.Product.GetCartItemByID(c.Request.Context(), uid, pid)
	if err != nil {
		models.ErrorResponse(c, http.StatusInternalServerError, "Internal service error", err.Error())
		return
	}

	c.JSON(http.StatusOK, models.JSON_Response{
		Success: true,
		Message: "Get product from cart successfully",
		Result:  result,
	})
}

// GetFullCart godoc
// @Summary      Get user's full cart
// @Description  Mendapatkan semua item keranjang milik user
// @Tags         cart
// @Accept       json
// @Accept       x-www-form-urlencoded
// @Produce      json
// @Security     ApiKeyAuth
// @Success      200  {object}  models.JSON_Response{result=[]models.CartItem} "Berhasil mendapatkan keranjang"
// @Failure      401  {object}  models.JSON_Response "Unauthorized"
// @Failure      500  {object}  models.JSON_Response "Internal server error"
// @Router       /cart [get]
func (pc *ProductsController) GetFullCart(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	if len(authHeader) < 8 {
		models.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized", "Invalid authorization header")
		return
	}

	claim, err := libs.VerifyJwt(authHeader[7:])
	if err != nil {
		models.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized", err.Error())
		return
	}

	id := int((*claim)["id"].(float64))

	results, err := pc.Product.GetCartByUserID(c.Request.Context(), id)
	if err != nil {
		models.ErrorResponse(c, http.StatusInternalServerError, "Server error", err.Error())
		return
	}

	c.JSON(http.StatusOK, models.JSON_Response{
		Success: true,
		Message: "Get all cart items successfully",
		Result:  results,
	})
}
