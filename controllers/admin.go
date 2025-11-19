package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/siddiq24/backend-coffee-shop/libs"
	"github.com/siddiq24/backend-coffee-shop/models"
)

type AdminController struct {
	Admin models.Admin
}

// ==================== PRODUCTS ====================

// GetAllProducts godoc
// @Summary Get all products
// @Description Mendapatkan daftar produk dengan paginasi dan filter
// @Tags admin
// @Accept json
// @Produce json
// @Param Page query int false "Nomor halaman" default(1)
// @Param Search query string false "Filter pencarian"
// @Security BearerAuth
// @Success 200 {object} models.JSON_Response
// @Router /admin/products [get]
func (a AdminController) GetAllProducts(c *gin.Context) {
	var req models.Param
	if err := c.ShouldBindQuery(&req); err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, "Invalid query parameters", err.Error())
		return
	}

	products, err := a.Admin.GetAllProducts(c, req.Page, req.Search)
	if err != nil {
		models.ErrorResponse(c, http.StatusInternalServerError, "Gagal mendapatkan products", err.Error())
		return
	}

	c.JSON(http.StatusOK, models.JSON_Response{
		Success: true,
		Message: "Berhasil mendapatkan products",
		Result:  products,
	})
}

// CreateProduct godoc
// @Summary Create product
// @Description Membuat produk baru
// @Tags admin
// @Accept json
// @Produce json
// @Param product body models.ProductCreateDTO true "Data produk"
// @Security BearerAuth
// @Success 201 {object} models.JSON_Response
// @Router /admin/products [post]
func (a AdminController) CreateProduct(c *gin.Context) {
	var req models.ProductCreateDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	err := a.Admin.CreateProduct(c.Request.Context(), req)
	if err != nil {
		models.ErrorResponse(c, http.StatusInternalServerError, "Gagal menambahkan product", err.Error())
		return
	}

	c.JSON(http.StatusCreated, models.JSON_Response{
		Success: true,
		Message: "Berhasil menambahkan product",
	})
}

// UpdateProduct godoc
// @Summary Update product
// @Description Mengupdate produk berdasarkan ID
// @Tags admin
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body models.ProductUpdateDTO true "Data produk"
// @Security BearerAuth
// @Success 200 {object} models.JSON_Response
// @Router /admin/products/{id} [patch]
func (a AdminController) UpdateProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, "Invalid product ID", err.Error())
		return
	}

	var req models.ProductUpdateDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	req.Id = id
	err = a.Admin.UpdateProduct(c.Request.Context(), req)
	if err != nil {
		models.ErrorResponse(c, http.StatusInternalServerError, "Gagal mengupdate product", err.Error())
		return
	}

	c.JSON(http.StatusOK, models.JSON_Response{
		Success: true,
		Message: "Berhasil mengupdate product",
	})
}

// DeleteProduct godoc
// @Summary Delete product
// @Description Menghapus produk (soft delete)
// @Tags admin
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Security BearerAuth
// @Success 200 {object} models.JSON_Response
// @Router /admin/products/{id} [delete]
func (a AdminController) DeleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, "Invalid product ID", err.Error())
		return
	}

	err = a.Admin.DeleteProduct(c.Request.Context(), id)
	if err != nil {
		models.ErrorResponse(c, http.StatusInternalServerError, "Gagal menghapus product", err.Error())
		return
	}

	c.JSON(http.StatusOK, models.JSON_Response{
		Success: true,
		Message: "Berhasil menghapus product",
	})
}

// ==================== CATEGORIES ====================

// GetAllCategories godoc
// @Summary Get all categories
// @Description Mendapatkan daftar kategori
// @Tags admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} models.JSON_Response
// @Router /admin/categories [get]
func (a AdminController) GetAllCategories(c *gin.Context) {
	categories, err := a.Admin.GetAllCategories(c.Request.Context())
	if err != nil {
		models.ErrorResponse(c, http.StatusInternalServerError, "Gagal mendapatkan categories", err.Error())
		return
	}

	c.JSON(http.StatusOK, models.JSON_Response{
		Success: true,
		Message: "Berhasil mendapatkan categories",
		Result:  categories,
	})
}

// CreateCategory godoc
// @Summary Create category
// @Description Membuat kategori baru
// @Tags admin
// @Accept json
// @Produce json
// @Param category body models.CategoryDTO true "Data kategori"
// @Security BearerAuth
// @Success 201 {object} models.JSON_Response
// @Router /admin/categories [post]
func (a AdminController) CreateCategory(c *gin.Context) {
	var req models.CategoryDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	err := a.Admin.CreateCategory(c.Request.Context(), req)
	if err != nil {
		models.ErrorResponse(c, http.StatusInternalServerError, "Gagal membuat category", err.Error())
		return
	}

	c.JSON(http.StatusCreated, models.JSON_Response{
		Success: true,
		Message: "Berhasil membuat category",
	})
}

// UpdateCategory godoc
// @Summary Update category
// @Description Mengupdate kategori
// @Tags admin
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param category body models.CategoryDTO true "Data kategori"
// @Security BearerAuth
// @Success 200 {object} models.JSON_Response
// @Router /admin/categories/{id} [patch]
func (a AdminController) UpdateCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, "Invalid category ID", err.Error())
		return
	}

	var req models.CategoryDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	req.Id = id
	err = a.Admin.UpdateCategory(c.Request.Context(), req)
	if err != nil {
		models.ErrorResponse(c, http.StatusInternalServerError, "Gagal mengupdate category", err.Error())
		return
	}

	c.JSON(http.StatusOK, models.JSON_Response{
		Success: true,
		Message: "Berhasil mengupdate category",
	})
}

// DeleteCategory godoc
// @Summary Delete category
// @Description Menghapus kategori
// @Tags admin
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Security BearerAuth
// @Success 200 {object} models.JSON_Response
// @Router /admin/categories/{id} [delete]
func (a AdminController) DeleteCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, "Invalid category ID", err.Error())
		return
	}

	err = a.Admin.DeleteCategory(c.Request.Context(), id)
	if err != nil {
		models.ErrorResponse(c, http.StatusInternalServerError, "Gagal menghapus category", err.Error())
		return
	}

	c.JSON(http.StatusOK, models.JSON_Response{
		Success: true,
		Message: "Berhasil menghapus category",
	})
}

// ==================== TRANSACTIONS ====================

// GetAllTransactions godoc
// @Summary Get all transactions
// @Description Mendapatkan daftar transaksi
// @Tags admin
// @Accept json
// @Produce json
// @Param Page query int false "Nomor halaman" default(1)
// @Param Status query string false "Filter status"
// @Security BearerAuth
// @Success 200 {object} models.JSON_Response
// @Router /admin/transactions [get]
func (a AdminController) GetAllTransactions(c *gin.Context) {
	var req models.TransactionParam
	if err := c.ShouldBindQuery(&req); err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, "Invalid query parameters", err.Error())
		return
	}

	transactions, err := a.Admin.GetAllTransactions(c.Request.Context(), req)
	if err != nil {
		models.ErrorResponse(c, http.StatusInternalServerError, "Gagal mendapatkan transactions", err.Error())
		return
	}

	c.JSON(http.StatusOK, models.JSON_Response{
		Success: true,
		Message: "Berhasil mendapatkan transactions",
		Result:  transactions,
	})
}

// CreateTransaction godoc
// @Summary Create transaction
// @Description Membuat transaksi baru
// @Tags admin
// @Accept json
// @Produce json
// @Param transaction body models.TransactionCreateDTO true "Data transaksi"
// @Security BearerAuth
// @Success 201 {object} models.JSON_Response
// @Router /admin/transactions [post]
func (a AdminController) CreateTransaction(c *gin.Context) {
	var req models.TransactionCreateDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	err := a.Admin.CreateTransaction(c.Request.Context(), req)
	if err != nil {
		models.ErrorResponse(c, http.StatusInternalServerError, "Gagal membuat transaction", err.Error())
		return
	}

	c.JSON(http.StatusCreated, models.JSON_Response{
		Success: true,
		Message: "Berhasil membuat transaction",
	})
}

// UpdateTransaction godoc
// @Summary Update transaction
// @Description Mengupdate transaksi (biasanya untuk status)
// @Tags admin
// @Accept json
// @Produce json
// @Param id path int true "Transaction ID"
// @Param transaction body models.TransactionUpdateDTO true "Data transaksi"
// @Security BearerAuth
// @Success 200 {object} models.JSON_Response
// @Router /admin/transactions/{id} [patch]
func (a AdminController) UpdateTransaction(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, "Invalid transaction ID", err.Error())
		return
	}

	var req models.TransactionUpdateDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	req.Id = id
	err = a.Admin.UpdateTransaction(c.Request.Context(), req)
	if err != nil {
		models.ErrorResponse(c, http.StatusInternalServerError, "Gagal mengupdate transaction", err.Error())
		return
	}

	c.JSON(http.StatusOK, models.JSON_Response{
		Success: true,
		Message: "Berhasil mengupdate transaction",
	})
}

// ==================== USERS ====================

// GetAllUsers godoc
// @Summary Get all users
// @Description Mendapatkan daftar users
// @Tags admin
// @Accept json
// @Produce json
// @Param Page query int false "Nomor halaman" default(1)
// @Param Search query string false "Filter pencarian"
// @Security BearerAuth
// @Success 200 {object} models.JSON_Response
// @Router /admin/users [get]
func (a AdminController) GetAllUsers(c *gin.Context) {
	var req models.Param
	if err := c.ShouldBindQuery(&req); err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, "Invalid query parameters", err.Error())
		return
	}

	users, err := a.Admin.GetAllUsers(c.Request.Context(), req)
	if err != nil {
		models.ErrorResponse(c, http.StatusInternalServerError, "Gagal mendapatkan users", err.Error())
		return
	}

	c.JSON(http.StatusOK, models.JSON_Response{
		Success: true,
		Message: "Berhasil mendapatkan users",
		Result:  users,
	})
}

// CreateUser godoc
// @Summary Create user
// @Description Membuat user baru
// @Tags admin
// @Accept json
// @Produce json
// @Param user body models.UserCreateDTO true "Data user"
// @Security BearerAuth
// @Success 201 {object} models.JSON_Response
// @Router /admin/users [post]
func (a AdminController) CreateUser(c *gin.Context) {
	var req models.UserCreateDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	err := a.Admin.CreateUser(c.Request.Context(), req)
	if err != nil {
		models.ErrorResponse(c, http.StatusInternalServerError, "Gagal membuat user", err.Error())
		return
	}

	c.JSON(http.StatusCreated, models.JSON_Response{
		Success: true,
		Message: "Berhasil membuat user",
	})
}

// UpdateUser godoc
// @Summary Update user
// @Description Mengupdate user
// @Tags admin
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body models.UserUpdateDTO true "Data user"
// @Security BearerAuth
// @Success 200 {object} models.JSON_Response
// @Router /admin/users/{id} [patch]
func (a AdminController) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, "Invalid user ID", err.Error())
		return
	}

	var req models.UserUpdateDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	req.Id = id
	err = a.Admin.UpdateUser(c.Request.Context(), req)
	if err != nil {
		models.ErrorResponse(c, http.StatusInternalServerError, "Gagal mengupdate user", err.Error())
		return
	}

	c.JSON(http.StatusOK, models.JSON_Response{
		Success: true,
		Message: "Berhasil mengupdate user",
	})
}

// DeleteUser godoc
// @Summary Delete user
// @Description Menghapus user (soft delete)
// @Tags admin
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Security BearerAuth
// @Success 200 {object} models.JSON_Response
// @Router /admin/users/{id} [delete]
func (a AdminController) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, "Invalid user ID", err.Error())
		return
	}

	err = a.Admin.DeleteUser(c.Request.Context(), id)
	if err != nil {
		models.ErrorResponse(c, http.StatusInternalServerError, "Gagal menghapus user", err.Error())
		return
	}

	c.JSON(http.StatusOK, models.JSON_Response{
		Success: true,
		Message: "Berhasil menghapus user",
	})
}

// ==================== PRODUCT IMAGES ====================

// GetProductImage godoc
// @Summary Get product image
// @Description Mendapatkan gambar produk berdasarkan ID
// @Tags admin
// @Accept json
// @Produce json
// @Param product_id path int true "Product ID"
// @Param image_id path int true "Image ID"
// @Security BearerAuth
// @Success 200 {object} models.JSON_Response
// @Router /admin/products/{product_id}/images/{image_id} [get]
func (a AdminController) GetProductImage(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("product_id"))
	if err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, "Invalid product ID", err.Error())
		return
	}

	imageID, err := strconv.Atoi(c.Param("image_id"))
	if err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, "Invalid image ID", err.Error())
		return
	}

	image, err := a.Admin.GetProductImage(c.Request.Context(), productID, imageID)
	if err != nil {
		models.ErrorResponse(c, http.StatusInternalServerError, "Gagal mendapatkan image", err.Error())
		return
	}

	c.JSON(http.StatusOK, models.JSON_Response{
		Success: true,
		Message: "Berhasil mendapatkan image",
		Result:  image,
	})
}

// AddProductImage godoc
// @Summary Add product image
// @Description Menambahkan gambar ke produk
// @Tags admin
// @Accept multipart/form-data
// @Produce json
// @Param product_id path int true "Product ID"
// @Param image formData file true "File gambar"
// @Security BearerAuth
// @Success 201 {object} models.JSON_Response
// @Router /admin/products/{product_id}/images [post]
func (a AdminController) AddProductImage(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("product_id"))
	if err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, "Invalid product ID", err.Error())
		return
	}

	// Ambil file dari form-data
	file, err := c.FormFile("image")
	if err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, "Gagal mengambil file", err.Error())
		return
	}

	// Simpan file ke folder lokal
	savedPath, err := libs.SaveUploadedFile(c, file, "images/products")
	if err != nil {
		models.ErrorResponse(c, http.StatusInternalServerError, "Gagal menyimpan file lokal", err.Error())
		return
	}

	// Upload ke cloudinary
	imageURL, err := libs.UploadToCloudinary(savedPath)
	if err != nil {
		models.ErrorResponse(c, http.StatusInternalServerError, "Gagal upload ke Cloudinary", err.Error())
		return
	}

	// Data untuk ke service
	req := models.ProductImageDTO{
		ProductId: productID,
		Image:     imageURL,
	}

	// Simpan ke database
	err = a.Admin.AddProductImage(c.Request.Context(), req)
	if err != nil {
		models.ErrorResponse(c, http.StatusInternalServerError, "Gagal menambahkan image", err.Error())
		return
	}

	c.JSON(http.StatusCreated, models.JSON_Response{
		Success: true,
		Message: "Berhasil menambahkan image",
		Result:  imageURL,
	})
}

// UpdateProductImage godoc
// @Summary Update product image
// @Description Mengupdate gambar produk
// @Tags admin
// @Accept multipart/form-data
// @Produce json
// @Param product_id path int true "Product ID"
// @Param image_id path int true "Image ID"
// @Param image formData file true "File gambar"
// @Security BearerAuth
// @Success 200 {object} models.JSON_Response
// @Router /admin/products/{product_id}/images/{image_id} [patch]
func (a AdminController) UpdateProductImage(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("product_id"))
	if err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, "Invalid product ID", err.Error())
		return
	}

	imageID, err := strconv.Atoi(c.Param("image_id"))
	if err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, "Invalid image ID", err.Error())
		return
	}

	// Ambil file upload baru
	file, err := c.FormFile("image")
	if err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, "Gagal mengambil file", err.Error())
		return
	}

	// Simpan file ke lokal
	savedPath, err := libs.SaveUploadedFile(c, file, "images/products")
	if err != nil {
		models.ErrorResponse(c, http.StatusInternalServerError, "Gagal menyimpan file lokal", err.Error())
		return
	}

	// Upload ke cloudinary
	imageURL, err := libs.UploadToCloudinary(savedPath)
	if err != nil {
		models.ErrorResponse(c, http.StatusInternalServerError, "Gagal upload ke Cloudinary", err.Error())
		return
	}

	req := models.ProductImageDTO{
		Id:        imageID,
		ProductId: productID,
		Image:     imageURL,
	}

	err = a.Admin.UpdateProductImage(c.Request.Context(), req)
	if err != nil {
		models.ErrorResponse(c, http.StatusInternalServerError, "Gagal update image", err.Error())
		return
	}

	c.JSON(http.StatusOK, models.JSON_Response{
		Success: true,
		Message: "Update image berhasil",
		Result:  imageURL,
	})
}

// DeleteProductImage godoc
// @Summary Delete product image
// @Description Menghapus gambar produk
// @Tags admin
// @Accept json
// @Produce json
// @Param product_id path int true "Product ID"
// @Param image_id path int true "Image ID"
// @Security BearerAuth
// @Success 200 {object} models.JSON_Response
// @Router /admin/products/{product_id}/images/{image_id} [delete]
func (a AdminController) DeleteProductImage(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("product_id"))
	if err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, "Invalid product ID", err.Error())
		return
	}

	imageID, err := strconv.Atoi(c.Param("image_id"))
	if err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, "Invalid image ID", err.Error())
		return
	}

	err = a.Admin.DeleteProductImage(c.Request.Context(), productID, imageID)
	if err != nil {
		models.ErrorResponse(c, http.StatusInternalServerError, "Gagal menghapus image", err.Error())
		return
	}

	c.JSON(http.StatusOK, models.JSON_Response{
		Success: true,
		Message: "Berhasil menghapus image",
	})
}

// SetFavProducts godoc
// @Summary Set favorite product
// @Description Menambahkan produk ke favorit
// @Tags admin
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Security BearerAuth
// @Success 200 {object} models.JSON_Response
// @Router /admin/products/{id}/favorite [post]
func (a AdminController) SetFavProducts(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, "Invalid Product ID", err.Error())
		return
	}

	if err := a.Admin.AddFavoriteProduct(c, id); err != nil {
		models.ErrorResponse(c, http.StatusInternalServerError, "Gagal menambahkan ke favorit", err.Error())
		return
	}

	c.JSON(http.StatusOK, models.JSON_Response{
		Success: true,
		Message: "Berhasil menambahkan product ke favorite",
	})
}
