package models

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type Admin struct{}

// ==================== STRUCTS & DTOs ====================

// Products
type ProductCreateDTO struct {
	Title       string  `json:"title" binding:"required"`
	Description string  `json:"description"`
	BasePrice   float64 `json:"base_price" binding:"required"`
	Stock       int     `json:"stock" binding:"required"`
	CategoryId  int     `json:"category_id" binding:"required"`
	Images      []Image `json:"images"`
	Sizes       []int   `json:"sizes"`
}

type ProductUpdateDTO struct {
	Id          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	BasePrice   float64 `json:"base_price"`
	Stock       int     `json:"stock"`
	CategoryId  int     `json:"category_id"`
}

type ProductListItem struct {
	Id           int     `json:"id"`
	Title        string  `json:"title"`
	Description  string  `json:"description"`
	BasePrice    float64 `json:"base_price"`
	Image        string  `json:"image"`
	Sizes        string  `json:"sizes"`
	Stock        int     `json:"stock"`
	CategoryName string  `json:"category_name"`
}

// Categories
type CategoryDTO struct {
	Id   int    `json:"id"`
	Name string `json:"name" binding:"required"`
}

// Transactions
type TransactionParam struct {
	Page   int    `form:"page"`
	Status string `form:"status"`
}

type TransactionCreateDTO struct {
	UserId     int               `json:"user_id" binding:"required"`
	ShippingId int               `json:"shipping_id" binding:"required"`
	TotalOrder float64           `json:"total_order" binding:"required"`
	StatusId   int               `json:"status_id"`
	PromoId    *int              `json:"promo_id"`
	Products   []OrderProductDTO `json:"products" binding:"required"`
}

type TransactionUpdateDTO struct {
	Id       int `json:"id"`
	StatusId int `json:"status_id" binding:"required"`
}

type TransactionListItem struct {
	Id         int       `json:"id"`
	NoOrder    string    `json:"no_order"`
	UserEmail  string    `json:"user_email"`
	TotalOrder float64   `json:"total_order"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
}

type OrderProductDTO struct {
	ProductId int  `json:"product_id" binding:"required"`
	SizeId    int  `json:"size_id" binding:"required"`
	VarianId  *int `json:"varian_id"`
	Qty       int  `json:"qty" binding:"required"`
}

// Users
type UserCreateDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Role     string `json:"role"`
}

type UserUpdateDTO struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

type UserListItem struct {
	Id        int       `json:"id"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

// Product Images
type ProductImageDTO struct {
	Id        int    `json:"id"`
	ProductId int    `json:"product_id"`
	Image     string `json:"image" binding:"required"`
}

type Image struct {
	Id  int    `json:"id"`
	Img string `json:"image"`
}

// Query Params
type Param struct {
	Page   int    `form:"page"`
	Search string `form:"search"`
}

// ==================== PRODUCTS METHODS ====================

func (a *Admin) GetAllProducts(ctx *gin.Context, page int, search string) ([]ProductListItem, error) {
	c := ctx.Request.Context()
	if page < 1 {
		page = 1
	}
	limit := 10
	offset := (page - 1) * limit
	searchPattern := "%" + search + "%"

	query := `
		SELECT 
			p.id,
			p.title,
			p.description,
			p.base_price,
			COALESCE(MAX(i.image), '') as image,
			COALESCE(STRING_AGG(DISTINCT sz.name, ', '), '') AS sizes,
			p.stock,
			COALESCE(c.name, '') as category_name
		FROM products p
		LEFT JOIN categories c ON c.id = p.category_id
		LEFT JOIN products_images i ON i.product_id = p.id
		LEFT JOIN products_sizes ps ON ps.product_id = p.id
		LEFT JOIN sizes sz ON sz.id = ps.size_id
		WHERE p.deleted_at IS NULL 
		AND (p.title ILIKE $3 OR p.description ILIKE $3)
		GROUP BY p.id, c.name
		ORDER BY p.id DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := Pg.Query(c, query, limit, offset, searchPattern)
	if err != nil {
		return nil, fmt.Errorf("failed to query products: %w", err)
	}
	defer rows.Close()

	var products []ProductListItem
	for rows.Next() {
		var p ProductListItem
		if err := rows.Scan(&p.Id, &p.Title, &p.Description, &p.BasePrice, &p.Image, &p.Sizes, &p.Stock, &p.CategoryName); err != nil {
			return nil, fmt.Errorf("failed to scan product: %w", err)
		}
		products = append(products, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return products, nil
}

func (a *Admin) CreateProduct(c context.Context, p ProductCreateDTO) error {
	conn, err := Pg.Acquire(c)
	if err != nil {
		return fmt.Errorf("failed to acquire connection: %w", err)
	}
	defer conn.Release()

	tx, err := conn.Begin(c)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback(c)
		}
	}()

	var productID int
	err = tx.QueryRow(c, `
		INSERT INTO products (title, description, base_price, stock, category_id) 
		VALUES ($1, $2, $3, $4, $5) 
		RETURNING id`,
		p.Title, p.Description, p.BasePrice, p.Stock, p.CategoryId,
	).Scan(&productID)
	if err != nil {
		return fmt.Errorf("failed to insert product: %w", err)
	}

	// Insert images
	for _, img := range p.Images {
		_, err = tx.Exec(c, `INSERT INTO products_images (product_id, image) VALUES ($1, $2)`, productID, img.Img)
		if err != nil {
			return fmt.Errorf("failed to insert product image: %w", err)
		}
	}

	// Insert sizes
	for _, sizeID := range p.Sizes {
		_, err = tx.Exec(c, `INSERT INTO products_sizes (product_id, size_id) VALUES ($1, $2)`, productID, sizeID)
		if err != nil {
			return fmt.Errorf("failed to insert product size: %w", err)
		}
	}

	if err := tx.Commit(c); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (a *Admin) UpdateProduct(c context.Context, p ProductUpdateDTO) error {
	query := `
		UPDATE products 
		SET title = COALESCE(NULLIF($2, ''), title),
		    description = COALESCE(NULLIF($3, ''), description),
		    base_price = COALESCE(NULLIF($4, 0), base_price),
		    stock = COALESCE(NULLIF($5, 0), stock),
		    category_id = COALESCE(NULLIF($6, 0), category_id),
		    updated_at = now()
		WHERE id = $1 AND deleted_at IS NULL
	`

	result, err := Pg.Exec(c, query, p.Id, p.Title, p.Description, p.BasePrice, p.Stock, p.CategoryId)
	if err != nil {
		return fmt.Errorf("failed to update product: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("product not found or already deleted")
	}

	return nil
}

func (a *Admin) DeleteProduct(c context.Context, id int) error {
	query := `UPDATE products SET deleted_at = now() WHERE id = $1`

	result, err := Pg.Exec(c, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete product: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("product not found")
	}

	return nil
}

// ==================== CATEGORIES METHODS ====================

func (a *Admin) GetAllCategories(c context.Context) ([]CategoryDTO, error) {
	query := `SELECT id, name FROM categories ORDER BY id ASC`
	rows, err := Pg.Query(c, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query categories: %w", err)
	}
	defer rows.Close()

	var categories []CategoryDTO
	for rows.Next() {
		var cat CategoryDTO
		if err := rows.Scan(&cat.Id, &cat.Name); err != nil {
			return nil, fmt.Errorf("failed to scan category: %w", err)
		}
		categories = append(categories, cat)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return categories, nil
}

func (a *Admin) CreateCategory(c context.Context, cat CategoryDTO) error {
	query := `INSERT INTO categories (name) VALUES ($1)`
	_, err := Pg.Exec(c, query, cat.Name)
	if err != nil {
		return fmt.Errorf("failed to create category: %w", err)
	}
	return nil
}

func (a *Admin) UpdateCategory(c context.Context, cat CategoryDTO) error {
	query := `UPDATE categories SET name = $2 WHERE id = $1`

	result, err := Pg.Exec(c, query, cat.Id, cat.Name)
	if err != nil {
		return fmt.Errorf("failed to update category: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("category not found")
	}

	return nil
}

func (a *Admin) DeleteCategory(c context.Context, id int) error {
	// Check if category is being used by any product
	var count int
	err := Pg.QueryRow(c, `SELECT COUNT(*) FROM products WHERE category_id = $1 AND deleted_at IS NULL`, id).Scan(&count)
	if err != nil {
		return fmt.Errorf("failed to check category usage: %w", err)
	}

	if count > 0 {
		return fmt.Errorf("cannot delete category: it is being used by %d products", count)
	}

	query := `DELETE FROM categories WHERE id = $1`

	result, err := Pg.Exec(c, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete category: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("category not found")
	}

	return nil
}

// ==================== TRANSACTIONS METHODS ====================

func (a *Admin) GetAllTransactions(c context.Context, param TransactionParam) ([]TransactionListItem, error) {
	if param.Page < 1 {
		param.Page = 1
	}
	limit := 10
	offset := (param.Page - 1) * limit

	query := `
		SELECT 
			o.id,
			o.no_order,
			u.email,
			o.total_order,
			s.name as status,
			o.created_at
		FROM orders o
		JOIN users u ON u.id = o.user_id
		JOIN status s ON s.id = o.status_id
		WHERE ($3 = '' OR s.name ILIKE '%' || $3 || '%')
		ORDER BY o.created_at DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := Pg.Query(c, query, limit, offset, param.Status)
	if err != nil {
		return nil, fmt.Errorf("failed to query transactions: %w", err)
	}
	defer rows.Close()

	var transactions []TransactionListItem
	for rows.Next() {
		var t TransactionListItem
		if err := rows.Scan(&t.Id, &t.NoOrder, &t.UserEmail, &t.TotalOrder, &t.Status, &t.CreatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan transaction: %w", err)
		}
		transactions = append(transactions, t)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return transactions, nil
}

func (a *Admin) CreateTransaction(c context.Context, t TransactionCreateDTO) error {
	conn, err := Pg.Acquire(c)
	if err != nil {
		return fmt.Errorf("failed to acquire connection: %w", err)
	}
	defer conn.Release()

	tx, err := conn.Begin(c)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback(c)
		}
	}()

	// Generate order number
	noOrder := fmt.Sprintf("ORD-%d-%d", time.Now().Unix(), t.UserId)

	var orderID int
	err = tx.QueryRow(c, `
		INSERT INTO orders (user_id, shipping_id, total_order, no_order, status_id, promo_id) 
		VALUES ($1, $2, $3, $4, $5, $6) 
		RETURNING id`,
		t.UserId, t.ShippingId, t.TotalOrder, noOrder, t.StatusId, t.PromoId,
	).Scan(&orderID)
	if err != nil {
		return fmt.Errorf("failed to insert order: %w", err)
	}

	// Insert order products
	for _, prod := range t.Products {
		_, err = tx.Exec(c, `
			INSERT INTO orders_products (order_id, product_id, size_id, varian_id, qty) 
			VALUES ($1, $2, $3, $4, $5)`,
			orderID, prod.ProductId, prod.SizeId, prod.VarianId, prod.Qty,
		)
		if err != nil {
			return fmt.Errorf("failed to insert order product: %w", err)
		}
	}

	if err := tx.Commit(c); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (a *Admin) UpdateTransaction(c context.Context, t TransactionUpdateDTO) error {
	query := `UPDATE orders SET status_id = $2, updated_at = now() WHERE id = $1`

	result, err := Pg.Exec(c, query, t.Id, t.StatusId)
	if err != nil {
		return fmt.Errorf("failed to update transaction: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("transaction not found")
	}

	return nil
}

// ==================== USERS METHODS ====================

func (a *Admin) GetAllUsers(c context.Context, param Param) ([]UserListItem, error) {
	if param.Page < 1 {
		param.Page = 1
	}
	limit := 10
	offset := (param.Page - 1) * limit
	searchPattern := "%" + param.Search + "%"

	query := `
		SELECT id, email, role, created_at 
		FROM users 
		WHERE deleted_at IS NULL 
		AND email ILIKE $3
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := Pg.Query(c, query, limit, offset, searchPattern)
	if err != nil {
		return nil, fmt.Errorf("failed to query users: %w", err)
	}
	defer rows.Close()

	var users []UserListItem
	for rows.Next() {
		var u UserListItem
		if err := rows.Scan(&u.Id, &u.Email, &u.Role, &u.CreatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return users, nil
}

func (a *Admin) CreateUser(c context.Context, u UserCreateDTO) error {
	if u.Role == "" {
		u.Role = "user"
	}

	// Hash password before storing (you should implement this)
	// hashedPassword, err := libs.HashPassword(u.Password)
	// if err != nil {
	//     return fmt.Errorf("failed to hash password: %w", err)
	// }

	query := `INSERT INTO users (email, password, role) VALUES ($1, $2, $3)`
	_, err := Pg.Exec(c, query, u.Email, u.Password, u.Role) // Use hashedPassword instead of u.Password
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func (a *Admin) UpdateUser(c context.Context, u UserUpdateDTO) error {
	query := `
		UPDATE users 
		SET email = COALESCE(NULLIF($2, ''), email),
		    role = COALESCE(NULLIF($3, ''), role)
		WHERE id = $1 AND deleted_at IS NULL
	`

	result, err := Pg.Exec(c, query, u.Id, u.Email, u.Role)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("user not found or already deleted")
	}

	return nil
}

func (a *Admin) DeleteUser(c context.Context, id int) error {
	query := `UPDATE users SET deleted_at = now() WHERE id = $1`

	result, err := Pg.Exec(c, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}

// ==================== PRODUCT IMAGES METHODS ====================

func (a *Admin) GetProductImage(c context.Context, productID, imageID int) (*ProductImageDTO, error) {
	query := `SELECT id, product_id, image FROM products_images WHERE id = $1 AND product_id = $2`

	var img ProductImageDTO
	err := Pg.QueryRow(c, query, imageID, productID).Scan(&img.Id, &img.ProductId, &img.Image)
	if err != nil {
		return nil, fmt.Errorf("failed to get product image: %w", err)
	}

	return &img, nil
}

func (a *Admin) AddProductImage(c context.Context, img ProductImageDTO) error {
	query := `INSERT INTO products_images (product_id, image) VALUES ($1, $2)`
	_, err := Pg.Exec(c, query, img.ProductId, img.Image)
	if err != nil {
		return fmt.Errorf("failed to add product image: %w", err)
	}
	return nil
}

func (a *Admin) UpdateProductImage(c context.Context, img ProductImageDTO) error {
	query := `UPDATE products_images SET image = $3 WHERE id = $1 AND product_id = $2`

	result, err := Pg.Exec(c, query, img.Id, img.ProductId, img.Image)
	if err != nil {
		return fmt.Errorf("failed to update product image: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("product image not found")
	}

	return nil
}

func (a *Admin) DeleteProductImage(c context.Context, productID, imageID int) error {
	query := `DELETE FROM products_images WHERE id = $1 AND product_id = $2`

	result, err := Pg.Exec(c, query, imageID, productID)
	if err != nil {
		return fmt.Errorf("failed to delete product image: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("product image not found")
	}

	return nil
}

// ==================== FAVORITE PRODUCTS METHODS ====================

func (a *Admin) AddFavoriteProduct(ctx *gin.Context, productID int) error {
	c := ctx.Request.Context()

	// Get user ID from context (assuming it's set by authentication middleware)
	userID, exists := ctx.Get("userID")
	if !exists {
		return fmt.Errorf("user not authenticated")
	}

	query := `INSERT INTO favorite_products (user_id, product_id) VALUES ($1, $2) ON CONFLICT DO NOTHING`

	result, err := Pg.Exec(c, query, userID, productID)
	if err != nil {
		return fmt.Errorf("failed to add favorite product: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("product already in favorites or not found")
	}

	return nil
}
