package models

import (
	"context"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
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

type Product_ress struct {
	Id        int             `json:"id"`
	Title     string          `json:"title"`
	Desc      string          `json:"desc"`
	Price     uint64          `json:"price"`
	Discount  sql.NullFloat64 `json:"discount"`
	Category  string          `json:"category"`
	Images    []string        `json:"images"`
	Sizes     []string        `json:"sizes"`
	Frequency int             `json:"freq,omitempty"`
}

type Product struct{}

func (p *Product) AllProductFiltered(c context.Context, prm Product_Params) ([]Product_ress, error) {
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
			"COALESCE(ARRAY_AGG(DISTINCT sz.name) FILTER (WHERE sz.name IS NOT NULL), '{}') AS sizes",
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
		var images, sizes []string

		if err := rows.Scan(
			&pr.Id,
			&pr.Title,
			&pr.Desc,
			&pr.Price,
			&pr.Discount,
			&pr.Category,
			&images,
			&sizes,
		); err != nil {
			return nil, fmt.Errorf("scan error: %w", err)
		}

		pr.Images = images
		pr.Sizes = sizes
		products = append(products, pr)
	}

	return products, nil
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
			"COALESCE(ARRAY_AGG(DISTINCT sz.name) FILTER (WHERE sz.name IS NOT NULL), '{}') AS sizes",
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
		var images, sizes []string

		if err := rows.Scan(
			&pr.Id,
			&pr.Title,
			&pr.Desc,
			&pr.Price,
			&pr.Discount,
			&pr.Category,
			&images,
			&sizes,
		); err != nil {
			return nil, fmt.Errorf("scan error: %w", err)
		}

		pr.Images = images
		pr.Sizes = sizes
		products = append(products, pr)
	}

	return products, nil
}

func (p *Product) GetRecommendation(ctx context.Context, id int, limit int) ([]Product_ress, error) {
	var products []Product_ress

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	query := psql.
		Select(
			"p2.id",
			"p2.title",
			"p2.description",
			"p2.base_price",
			"COALESCE(pr.discount, 0) AS discount",
			"c.name AS category_name",
			"COALESCE(ARRAY_AGG(DISTINCT i.image) FILTER (WHERE i.image IS NOT NULL), '{}') AS images",
			"COALESCE(ARRAY_AGG(DISTINCT sz.name) FILTER (WHERE sz.name IS NOT NULL), '{}') AS sizes",
			"COUNT(*) AS frequency",
		).
		From("products p1").
		JoinClause("FULL JOIN products_tags pt ON pt.product_id = p1.id").
		JoinClause("FULL JOIN products_tags pt2 ON pt2.tag_id = pt.tag_id").
		JoinClause("FULL JOIN orders_products op ON op.product_id = p1.id").
		JoinClause("FULL JOIN orders_products op2 ON op2.order_id = op.order_id").
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
		Limit(uint64(limit))

	sqlStr, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building query: %w", err)
	}

	rows, err := Pg.Query(ctx, sqlStr, args...)
	if err != nil {
		return nil, fmt.Errorf("query error: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var pr Product_ress
		var images, sizes []string

		if err := rows.Scan(
			&pr.Id,
			&pr.Title,
			&pr.Desc,
			&pr.Price,
			&pr.Discount,
			&pr.Category,
			&images,
			&sizes,
			&pr.Frequency,
		); err != nil {
			return nil, fmt.Errorf("scan error: %w", err)
		}

		pr.Images = images
		pr.Sizes = sizes
		products = append(products, pr)
	}

	return products, nil
}

type Cart_Request struct {
	OrderId   int `json:"order_id" form:"order_id"`
	UserId    int `json:"user_id" form:"user_id"`
	ProductId int `json:"product_id" form:"product_id"`
	VarianId  int `json:"varian_id" form:"varian_id"`
	SizeId    int `json:"size_id"`
	Qty       int `json:"quantity" form:"quantity"`
}

func (p *Product) CreateCart(c context.Context, req Cart_Request) (Cart_Request, error) {
	var order_id uint64
	if err := Pg.QueryRow(c, `INSERT INTO orders(user_id) VALUES ($1) RETURNING id `, req.UserId).Scan(&order_id); err != nil {
		return Cart_Request{}, fmt.Errorf("failed insert product: %w", err)
	}
	if _, err := Pg.Exec(c, `INSERT INTO orderss_products(order_id, product_id, varian_id, size_id, qty) VALUES ($1, $2, $3, $4, $5)`); err != nil {
		return Cart_Request{}, fmt.Errorf("failed insert cart %w", err)
	}
	if err := Pg.QueryRow(c, `
		SELECT 
			p.id,
			s.name,
			v.name
		FROM orders_products op
		LEFT JOIN products p ON p.id = op.product_id
		LEFT JOIN sizes s ON s.id = op.size_id
		LEFT JOIN variants v ON v.id = op.varian_id
		WHERE op.order_id = $1
	`, req.UserId).Scan(&req.OrderId, &req.ProductId, &req.SizeId, &req.VarianId, &req.Qty); err != nil {
		return Cart_Request{}, fmt.Errorf("failed insert to products %w", err)
	}

	return req, nil
}

type ProductCart struct {
	Id       int
	Image    string
	Title    string
	Qty      int
	Size     string
	Variants string
}

func (p *Product) GetProductCart(c context.Context, id int) ([]ProductCart, error) {
	rows, err := Pg.Query(c, `
		SELECT 
			p.id,
			p.title,
			pi.image,
			s.name,
			v.name,
			op.qty,
			MAX(op.order_id)
		FROM orders_products op
		FULL JOIN products p ON p.id = op.product_id
		FULL JOIN sizes s ON s.id = op.size_id
		FULL JOIN variants v ON v.id = op.varian_id
		full JOIN products_images pi ON pi.product_id = p.id
		left join orders o On o.id = op.order_id
		WHERE o.user_id = $1
		GROUP BY p.id, pi.id, s.id, v.id, op.qty;;
	`, id)
	if err != nil {
		return nil, fmt.Errorf("Error binding query %w", err)
	}

	var ress []ProductCart
	defer rows.Close()

	for rows.Next() {
		rsp := ProductCart{}
		err := rows.Scan(
			&rsp.Id,
			&rsp.Title,
			&rsp.Image,
			&rsp.Size,
			&rsp.Variants,
		)
		if err != nil {
			return nil, fmt.Errorf("Scan error %w", err)
		}
		ress = append(ress, rsp)
	}
	return ress, nil
}
