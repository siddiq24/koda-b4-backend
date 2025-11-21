package models

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"math"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

type Product_Params struct {
	Page       int      `form:"page"`
	Search     string   `form:"search"`
	CategoryId []string `form:"cat"`
	MinPrice   uint64   `form:"minPrice"`
	MaxPrice   uint64   `form:"maxPrice"`
	ShortBy    string   `form:"shortBy"`
	Asc        bool     `form:"asc"`
	Limit      int      `form:"limit"`
}

type ProductsRessponse struct {
	Success    bool   `json:"success"`
	Message    string `json:"message"`
	Page       int    `json:"page"`
	NextPage   string `json:"next_page"`
	PrevPage   string `json:"prev_page"`
	TotalPages int    `json:"total_page"`
	Result     any    `json:"result,omitempty"`
}

type Product_ress struct {
	Id        int             `json:"id"`
	Title     string          `json:"title"`
	Desc      string          `json:"desc"`
	Price     uint64          `json:"price"`
	Stock     uint64          `json:"stock"`
	Discount  sql.NullFloat64 `json:"discount"`
	Category  string          `json:"category"`
	Images    []string        `json:"images"`
	Sizes     []Size          `json:"sizes"`
	Variants  []Variant       `json:"variants"`
	Frequency int             `json:"freq,omitempty"`
}

type Size struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
type Variant struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Product struct{}

func (p *Product) AllProductFiltered(c context.Context, prm Product_Params) ([]Product_ress, int, error) {
	var products []Product_ress
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	sortDir := "DESC"
	if prm.Asc {
		sortDir = "ASC"
	}

	query := psql.
		Select(
			"p.id",
			"p.title",
			"p.description",
			"p.base_price",
			"COALESCE(pr.discount, 0) AS discount",
			"c.name AS category_name",
			"COALESCE(ARRAY_AGG(DISTINCT i.image) FILTER (WHERE i.image IS NOT NULL), '{}') AS images",
			"COALESCE(json_agg(DISTINCT jsonb_build_object('id', sz.id, 'name', sz.name)) FILTER (WHERE sz.id IS NOT NULL), '[]') AS sizes",
		).
		From("products p").
		LeftJoin("categories c ON c.id = p.category_id").
		LeftJoin("products_images i ON i.product_id = p.id").
		LeftJoin("products_sizes ps ON ps.product_id = p.id").
		LeftJoin("sizes sz ON sz.id = ps.size_id").
		LeftJoin("products_promos pp ON pp.product_id = p.id").
		LeftJoin("promos pr ON pr.id = pp.promo_id").
		GroupBy("p.id", "c.name", "pr.discount")

	if prm.Search != "" {
		search := "%" + prm.Search + "%"
		query = query.Where(sq.Or{
			sq.ILike{"p.title": search},
			sq.ILike{"p.description": search},
		})
	}

	if len(prm.CategoryId) > 0 {
		query = query.Where(sq.Eq{"p.category_id": prm.CategoryId})
	}

	if prm.MinPrice > 0 {
		query = query.Where(sq.GtOrEq{"p.base_price": prm.MinPrice})
	}

	if prm.MaxPrice > 0 {
		query = query.Where(sq.LtOrEq{"p.base_price": prm.MaxPrice})
	}

	switch prm.ShortBy {
	case "title":
		query = query.OrderBy(fmt.Sprintf("p.title %s", sortDir))
	case "price":
		query = query.OrderBy(fmt.Sprintf("p.base_price %s", sortDir))
	default:
		query = query.OrderBy(fmt.Sprintf("p.id %s", sortDir))
	}

	if prm.Page <= 0 {
		prm.Page = 1
	}
	offset := (prm.Page - 1) * prm.Limit
	query = query.Limit(uint64(prm.Limit)).Offset(uint64(offset))

	countQuery := psql.
		Select("COUNT(DISTINCT p.id)").
		From("products p").
		LeftJoin("categories c ON c.id = p.category_id")

	if prm.Search != "" {
		search := "%" + prm.Search + "%"
		countQuery = countQuery.Where(sq.Or{
			sq.ILike{"p.title": search},
			sq.ILike{"p.description": search},
		})
	}
	if len(prm.CategoryId) > 0 {
		countQuery = countQuery.Where(sq.Eq{"p.category_id": prm.CategoryId})
	}
	if prm.MinPrice > 0 {
		countQuery = countQuery.Where(sq.GtOrEq{"p.base_price": prm.MinPrice})
	}
	if prm.MaxPrice > 0 {
		countQuery = countQuery.Where(sq.LtOrEq{"p.base_price": prm.MaxPrice})
	}

	countSQL, countArgs, _ := countQuery.ToSql()

	var totalRows int
	if err := Pg.QueryRow(c, countSQL, countArgs...).Scan(&totalRows); err != nil {
		return nil, 0, fmt.Errorf("count error: %w", err)
	}

	sqlStr, args, err := query.ToSql()
	if err != nil {
		return nil, 0, fmt.Errorf("failed to build SQL: %w", err)
	}

	rows, err := Pg.Query(c, sqlStr, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("query error: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var pr Product_ress
		var images []string
		var sizesJSON []byte

		if err := rows.Scan(
			&pr.Id,
			&pr.Title,
			&pr.Desc,
			&pr.Price,
			&pr.Discount,
			&pr.Category,
			&images,
			&sizesJSON,
		); err != nil {
			return nil, 0, fmt.Errorf("scan error: %w", err)
		}

		pr.Images = images

		var sizes []Size
		if err := json.Unmarshal(sizesJSON, &sizes); err != nil {
			return nil, 0, fmt.Errorf("failed to unmarshal sizes: %w", err)
		}
		pr.Sizes = sizes

		products = append(products, pr)
	}

	totalPages := int(math.Ceil(float64(totalRows) / float64(prm.Limit)))

	return products, totalPages, nil
}

func (p *Product) FavProducts(c context.Context, limit int) ([]Product_ress, error) {
	var products []Product_ress

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	query := psql.
		Select(
			"p.id",
			"p.title",
			"p.description",
			"p.base_price",
			"COALESCE(pr.discount, 0) AS discount",
			"c.name AS category_name",
			"COALESCE(ARRAY_AGG(DISTINCT i.image) FILTER (WHERE i.image IS NOT NULL), '{}') AS images",
			"COALESCE(json_agg(DISTINCT jsonb_build_object('id', sz.id, 'name', sz.name)) FILTER (WHERE sz.id IS NOT NULL), '[]') AS sizes",
		).
		From("products p").
		LeftJoin("categories c ON c.id = p.category_id").
		LeftJoin("products_images i ON i.product_id = p.id").
		LeftJoin("products_sizes ps ON ps.product_id = p.id").
		LeftJoin("sizes sz ON sz.id = ps.size_id").
		LeftJoin("products_promos pp ON pp.product_id = p.id").
		LeftJoin("promos pr ON pr.id = pp.promo_id").
		Where("p.is_favorite = true").
		GroupBy("p.id", "c.name", "pr.discount").
		OrderBy("p.id ASC").
		Limit(uint64(limit))

	sqlStr, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build SQL: %w", err)
	}

	rows, err := Pg.Query(c, sqlStr, args...)
	if err != nil {
		return nil, fmt.Errorf("query error: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var pr Product_ress
		var images []string
		var sizesJSON []byte

		if err := rows.Scan(
			&pr.Id,
			&pr.Title,
			&pr.Desc,
			&pr.Price,
			&pr.Discount,
			&pr.Category,
			&images,
			&sizesJSON,
		); err != nil {
			return nil, fmt.Errorf("scan error: %w", err)
		}

		pr.Images = images

		var sizes []Size
		if err := json.Unmarshal(sizesJSON, &sizes); err != nil {
			return nil, fmt.Errorf("failed to unmarshal sizes: %w", err)
		}
		pr.Sizes = sizes

		products = append(products, pr)
	}

	return products, nil
}

func (p *Product) GetProductByID(c context.Context, productID int64) (Product_ress, error) {
	var product Product_ress
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	query := psql.
		Select(
			"p.id",
			"p.title",
			"p.description",
			"p.base_price",
			"p.stock",
			"COALESCE(pr.discount, 0) AS discount",
			"c.name AS category_name",
			"COALESCE(ARRAY_AGG(DISTINCT i.image) FILTER (WHERE i.image IS NOT NULL), '{}') AS images",
			"COALESCE(json_agg(DISTINCT jsonb_build_object('id', sz.id, 'name', sz.name)) FILTER (WHERE sz.id IS NOT NULL), '[]') AS sizes",
			"COALESCE(json_agg(DISTINCT jsonb_build_object('id', v.id, 'name', v.name)) FILTER (WHERE v.id IS NOT NULL), '[]') AS variants",
		).
		From("products p").
		LeftJoin("categories c ON c.id = p.category_id").
		LeftJoin("products_images i ON i.product_id = p.id").
		LeftJoin("products_sizes ps ON ps.product_id = p.id").
		LeftJoin("sizes sz ON sz.id = ps.size_id").
		LeftJoin("products_variants pv ON pv.product_id = p.id").
		LeftJoin("variants v ON v.id = pv.variant_id").
		LeftJoin("products_promos pp ON pp.product_id = p.id").
		LeftJoin("promos pr ON pr.id = pp.promo_id").
		Where(sq.Eq{"p.id": productID}).
		GroupBy("p.id", "c.name", "pr.discount")

	sqlStr, args, err := query.ToSql()
	if err != nil {
		return product, fmt.Errorf("failed to build SQL: %w", err)
	}

	var images []string
	var sizesJSON []byte
	var variantJSON []byte

	err = Pg.QueryRow(c, sqlStr, args...).Scan(
		&product.Id,
		&product.Title,
		&product.Desc,
		&product.Price,
		&product.Stock,
		&product.Discount,
		&product.Category,
		&images,
		&sizesJSON,
		&variantJSON,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return product, fmt.Errorf("product not found")
		}
		return product, fmt.Errorf("query error: %w", err)
	}

	product.Images = images

	var sizes []Size
	if err := json.Unmarshal(sizesJSON, &sizes); err != nil {
		return product, fmt.Errorf("failed to unmarshal sizes: %w", err)
	}
	product.Sizes = sizes

	var variants []Variant
	if err := json.Unmarshal(variantJSON, &variants); err != nil {
		return product, fmt.Errorf("failed to unmarshal sizes: %w", err)
	}
	product.Variants = variants

	return product, nil
}

func (p *Product) GetRecommendation(c context.Context, id int, page, limit int) ([]Product_ress, int, error) {
	var products []Product_ress

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	offset := (page - 1) * limit

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	countQuery := psql.
		Select("COUNT(DISTINCT p2.id)").
		From("products p1").
		JoinClause("FULL JOIN products_tags pt ON pt.product_id = p1.id").
		JoinClause("FULL JOIN products_tags pt2 ON pt2.tag_id = pt.tag_id").
		JoinClause("FULL JOIN orders_products op ON op.product_id = p1.id").
		JoinClause("FULL JOIN orders_products op2 ON op2.invoice = op.invoice").
		JoinClause("FULL JOIN products p2 ON p2.id = pt2.product_id OR p2.id = op2.product_id OR p2.category_id = p1.category_id").
		Where(sq.And{
			sq.Eq{"p1.id": id},
			sq.NotEq{"p2.id": id},
		})

	countSqlStr, countArgs, err := countQuery.ToSql()
	if err != nil {
		return nil, 0, fmt.Errorf("error building count query: %w", err)
	}

	var totalCount int
	err = Pg.QueryRow(c, countSqlStr, countArgs...).Scan(&totalCount)
	if err != nil {
		return nil, 0, fmt.Errorf("count query error: %w", err)
	}

	query := psql.
		Select(
			"p2.id",
			"p2.title",
			"p2.description",
			"p2.base_price",
			"p2.stock",
			"COALESCE(pr.discount, 0) AS discount",
			"c.name AS category_name",
			"COALESCE(ARRAY_AGG(DISTINCT i.image) FILTER (WHERE i.image IS NOT NULL), '{}') AS images",
			"COALESCE(json_agg(DISTINCT jsonb_build_object('id', sz.id, 'name', sz.name)) FILTER (WHERE sz.id IS NOT NULL), '[]') AS sizes",
			"COUNT(*) AS frequency",
		).
		From("products p1").
		JoinClause("FULL JOIN products_tags pt ON pt.product_id = p1.id").
		JoinClause("FULL JOIN products_tags pt2 ON pt2.tag_id = pt.tag_id").
		JoinClause("FULL JOIN orders_products op ON op.product_id = p1.id").
		JoinClause("FULL JOIN orders_products op2 ON op2.invoice = op.invoice").
		JoinClause("FULL JOIN products p2 ON p2.id = pt2.product_id OR p2.id = op2.product_id OR p2.category_id = p1.category_id").
		LeftJoin("categories c ON c.id = p2.category_id").
		LeftJoin("products_images i ON i.product_id = p2.id").
		LeftJoin("products_sizes ps ON ps.product_id = p2.id").
		LeftJoin("sizes sz ON sz.id = ps.size_id").
		LeftJoin("products_promos pp ON pp.product_id = p2.id").
		LeftJoin("promos pr ON pr.id = pp.promo_id").
		Where(sq.And{
			sq.Eq{"p1.id": id},
			sq.NotEq{"p2.id": id},
		}).
		GroupBy("p2.id", "c.name", "pr.discount").
		OrderBy("frequency DESC").
		Limit(uint64(limit)).
		Offset(uint64(offset))

	sqlStr, args, err := query.ToSql()
	if err != nil {
		return nil, 0, fmt.Errorf("error building query: %w", err)
	}

	rows, err := Pg.Query(c, sqlStr, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("query error: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var pr Product_ress
		var images []string
		var sizesJSON []byte

		if err := rows.Scan(
			&pr.Id,
			&pr.Title,
			&pr.Desc,
			&pr.Price,
			&pr.Stock,
			&pr.Discount,
			&pr.Category,
			&images,
			&sizesJSON,
			&pr.Frequency,
		); err != nil {
			return nil, 0, fmt.Errorf("scan error: %w", err)
		}

		pr.Images = images

		var sizes []Size
		if err := json.Unmarshal(sizesJSON, &sizes); err != nil {
			return nil, 0, fmt.Errorf("failed to unmarshal sizes: %w", err)
		}
		pr.Sizes = sizes

		products = append(products, pr)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("rows iteration error: %w", err)
	}

	return products, totalCount, nil
}

type CartRequest struct {
	OrderId   int `json:"order_id" form:"order_id"`
	UserId    int `json:"user_id" form:"user_id"`
	ProductId int `json:"product_id" form:"product_id"`
	VarianId  int `json:"varian_id" form:"varian_id"`
	SizeId    int `json:"size_id" form:"size_id"`
	Qty       int `json:"quantity" form:"quantity"`
}

type CartItem struct {
	ID          int64   `json:"id"`
	UserId      int     `json:"user_id"`
	ProductId   int     `json:"product_id"`
	VarianId    *int    `json:"varian_id"`
	SizeId      *int    `json:"size_id"`
	Qty         int     `json:"quantity"`
	Subtotal    float64 `json:"subtotal"`
	ProductName string  `json:"product_name"`
}

func (r *Product) AddToCart(c context.Context, req CartRequest) (*CartItem, error) {
	var basePrice float64
	var productName string

	err := Pg.QueryRow(c,
		"SELECT base_price, title FROM products WHERE id = $1 AND deleted_at IS NULL",
		req.ProductId,
	).Scan(&basePrice, &productName)
	if err != nil {
		return nil, fmt.Errorf("product not found: %w", err)
	}

	subtotal := basePrice * float64(req.Qty)

	var existingCartID int64
	var existingQty int

	err = Pg.QueryRow(c, `
		SELECT id, qty FROM carts 
		WHERE user_id = $1 AND product_id = $2 
		AND (varian_id = $3 OR ($3 IS NULL AND varian_id IS NULL))
		AND (size_id = $4 OR ($4 IS NULL AND size_id IS NULL))
	`, req.UserId, req.ProductId, req.VarianId, req.SizeId).Scan(&existingCartID, &existingQty)

	if err != nil && err != pgx.ErrNoRows {
		return nil, fmt.Errorf("failed to check existing cart: %w", err)
	}

	if err == nil {
		newQty := existingQty + req.Qty
		newSubtotal := basePrice * float64(newQty)

		_, err = Pg.Exec(c, `
			UPDATE carts 
			SET qty = $1, subtotal = $2, product_name = $3
			WHERE id = $4
		`, newQty, newSubtotal, productName, existingCartID)
		if err != nil {
			return nil, fmt.Errorf("failed to update cart: %w", err)
		}

		return &CartItem{
			ID:          existingCartID,
			UserId:      req.UserId,
			ProductId:   req.ProductId,
			VarianId:    &req.VarianId,
			SizeId:      &req.SizeId,
			Qty:         newQty,
			Subtotal:    newSubtotal,
			ProductName: productName,
		}, nil
	}

	var cartID int64
	err = Pg.QueryRow(c, `
		INSERT INTO carts (user_id, product_id, varian_id, size_id, qty, subtotal, product_name)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id
	`, req.UserId, req.ProductId, req.VarianId, req.SizeId, req.Qty, subtotal, productName).Scan(&cartID)
	if err != nil {
		return nil, fmt.Errorf("failed to add to cart: %w", err)
	}

	return &CartItem{
		ID:          cartID,
		UserId:      req.UserId,
		ProductId:   req.ProductId,
		VarianId:    &req.VarianId,
		SizeId:      &req.SizeId,
		Qty:         req.Qty,
		Subtotal:    subtotal,
		ProductName: productName,
	}, nil
}

func (r *Product) GetCartByUserID(c context.Context, userID int) ([]CartItem, error) {
	query := `
		SELECT 
			c.id,
			c.user_id,
			c.product_id,
			c.varian_id,
			c.size_id,
			c.qty,
			c.subtotal,
			c.product_name,
			v.name as variant_name,
			s.name as size_name
		FROM carts c
		LEFT JOIN variants v ON c.varian_id = v.id
		LEFT JOIN sizes s ON c.size_id = s.id
		WHERE c.user_id = $1
		ORDER BY c.id DESC`

	rows, err := Pg.Query(c, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get cart: %w", err)
	}
	defer rows.Close()

	var cartItems []CartItem
	for rows.Next() {
		var item CartItem
		var variantName, sizeName sql.NullString

		err := rows.Scan(
			&item.ID,
			&item.UserId,
			&item.ProductId,
			&item.VarianId,
			&item.SizeId,
			&item.Qty,
			&item.Subtotal,
			&item.ProductName,
			&variantName,
			&sizeName,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan cart item: %w", err)
		}
		cartItems = append(cartItems, item)
	}

	return cartItems, nil
}

func (r *Product) UpdateCartItem(c context.Context, cartID int, qty int) error {
	var productID int
	var currentSubtotal float64
	var currentQty int

	err := Pg.QueryRow(c,
		"SELECT product_id, subtotal, qty FROM carts WHERE id = $1",
		cartID,
	).Scan(&productID, &currentSubtotal, &currentQty)
	if err != nil {
		return fmt.Errorf("cart item not found: %w", err)
	}

	var basePrice float64
	err = Pg.QueryRow(c,
		"SELECT base_price FROM products WHERE id = $1",
		productID,
	).Scan(&basePrice)
	if err != nil {
		return fmt.Errorf("product not found: %w", err)
	}

	newSubtotal := basePrice * float64(qty)

	_, err = Pg.Exec(c, `
		UPDATE carts 
		SET qty = $1, subtotal = $2 
		WHERE id = $3`,
		qty, newSubtotal, cartID,
	)
	if err != nil {
		return fmt.Errorf("failed to update cart item: %w", err)
	}

	return nil
}

func (r *Product) DeleteCartItem(c context.Context, cartID int) error {
	_, err := Pg.Exec(c, "DELETE FROM carts WHERE id = $1", cartID)
	if err != nil {
		return fmt.Errorf("failed to delete cart item: %w", err)
	}
	return nil
}

func (r *Product) ClearUserCart(c context.Context, userID int) error {
	_, err := Pg.Exec(c, "DELETE FROM carts WHERE user_id = $1", userID)
	if err != nil {
		return fmt.Errorf("failed to clear user cart: %w", err)
	}
	return nil
}

func (r *Product) GetCartItemByID(c context.Context, userID, cartID int) (*CartItem, error) {
	var item CartItem

	err := Pg.QueryRow(c, `
		SELECT 
			id, user_id, product_id, varian_id, size_id, qty, subtotal, product_name
		FROM carts 
		WHERE id = $1 AND user_id = $2`,
		cartID, userID,
	).Scan(
		&item.ID,
		&item.UserId,
		&item.ProductId,
		&item.VarianId,
		&item.SizeId,
		&item.Qty,
		&item.Subtotal,
		&item.ProductName,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get cart item: %w", err)
	}

	return &item, nil
}
