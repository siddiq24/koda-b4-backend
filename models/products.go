package models

import (
	"context"
	"database/sql"
	"fmt"
	"slices"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Product_Params struct {
	Page       int
	Search     string
	CategoryId int
	MinPrice   uint64
	MaxPrice   uint64
}

type Product_ress struct {
	Id       int             `json:"id"`
	Title    string          `json:"title"`
	Desc     string          `json:"desc"`
	Price    uint64          `json:"price"`
	Discount sql.NullFloat64 `json:"discount"`
	Category string          `json:"category"`
	Images   []string        `json:"images"`
	Sizes    []string        `json:"sizes"`
}

type Product struct {
	Pg *pgxpool.Pool
}

func (p *Product) AllProductFiltered(c context.Context, prm Product_Params) ([]Product_ress, error) {
	productMap := make(map[int]*Product_ress)
	var res []Product_ress
	Limit := 10
	if prm.Page <= 1 {
		prm.Page = 1
	}
	offset := (prm.Page - 1) * Limit
	var filters []string
	if prm.Search != "" {
		search := strings.ReplaceAll(prm.Search, "'", "''")
		filters = append(filters, fmt.Sprintf("(p.title ILIKE '%%%s%%' OR p.description ILIKE '%%%s%%')", search, search))
	}
	if prm.CategoryId > 0 {
		filters = append(filters, fmt.Sprintf("p.category_id = %d", prm.CategoryId))
	}
	if prm.MinPrice > 0 {
		filters = append(filters, fmt.Sprintf("p.base_price >= %d", prm.MinPrice))
	}
	if prm.MaxPrice > 0 {
		filters = append(filters, fmt.Sprintf("p.base_price <= %d", prm.MaxPrice))
	}

	filterQuery := ""
	if len(filters) > 0 {
		filterQuery = "WHERE " + strings.Join(filters, " AND ")
	}

	query := fmt.Sprintf(`
		SELECT 
			p.id,
			p.title,
			p.description,
			p.base_price,
			pr.discount,
			c.name AS category_name,
			i.image,
			sz.name AS size_name
		FROM products p
		LEFT JOIN categories c ON c.id = p.category_id
		LEFT JOIN products_images i ON i.product_id = p.id
		LEFT JOIN products_sizes ps ON ps.product_id = p.id
		LEFT JOIN sizes sz ON sz.id = ps.size_id
		LEFT JOIN products_promos pp ON pp.product_id = p.id
		LEFT JOIN promos pr ON pr.id = pp.promo_id
		%s
		ORDER BY p.id ASC
		LIMIT %d OFFSET %d
	`, filterQuery, Limit, offset)

	rows, err := p.Pg.Query(c, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id       int
			title    string
			desc     string
			price    uint64
			discount sql.NullFloat64
			category *string
			image    *string
			size     *string
		)

		if err := rows.Scan(&id, &title, &desc, &price, &discount, &category, &image, &size); err != nil {
			return nil, err
		}

		if _, exist := productMap[id]; !exist {
			productMap[id] = &Product_ress{
				Id:       id,
				Title:    title,
				Desc:     desc,
				Price:    price,
				Discount: discount,
				Category: "",
				Images:   []string{},
				Sizes:    []string{},
			}
			if category != nil {
				productMap[id].Category = *category
			}
		}

		if image != nil && !slices.Contains(productMap[id].Images, *image) {
			productMap[id].Images = append(productMap[id].Images, *image)
		}
		if size != nil && !slices.Contains(productMap[id].Sizes, *size) {
			productMap[id].Sizes = append(productMap[id].Sizes, *size)
		}
	}

	for _, v := range productMap {
		res = append(res, *v)
	}

	return res, nil
}
