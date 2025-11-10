package models

import (
	"github.com/gin-gonic/gin"
)

type Products_admin struct {
	Id        int    `json:"id" form:"id"`
	Image     string `json:"image" form:"image"`
	Title     string `json:"title" form:"title"`
	Price     uint64 `json:"price" form:"price"`
	Desc      string `json:"desc" form:"desc"`
	Sizes     string `json:"sizes" form:"sizes"`
	Shippings string `json:"shippings" form:"shippings"`
	Stock     int    `json:"stock" form:"stock"`
}

type Param struct {
	Page   int    `query:"page"`
	Search string `query:"search"`
}

type Admin struct{}

func (a *Admin) GetAllProducts(ctx *gin.Context, page int, s string) ([]Products_admin, error) {
	c := ctx.Request.Context()
	limit := 10
	offset := (page - 1) * limit
	search := "%" + s + "%"
	Query := `
		SELECT 
			p.id,
			p.title,
			p.description,
			p.base_price,
			MAX(i.image),
			STRING_AGG(sz.name, ', ') AS size,
			p.stock
		FROM products p
		LEFT JOIN categories c ON c.id = p.category_id
		LEFT JOIN products_images i ON i.product_id = p.id
		LEFT JOIN products_sizes ps ON ps.product_id = p.id
		LEFT JOIN sizes sz ON sz.id = ps.size_id
		LEFT JOIN products_promos pp ON pp.product_id = p.id
		LEFT JOIN promos pr ON pr.id = pp.promo_id
		WHERE p.title ILIKE $3 OR p.description ILIKE '%$3%'
		GROUP BY p.id, c.name
		ORDER BY p.id ASC
		LIMIT $1 OFFSET $2
	`
	rows, err := Pg.Query(c, Query, limit, offset, search)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Products_admin
	for rows.Next() {
		var (
			id    int
			title string
			desc  string
			price uint64
			img   string
			size  string
			stock int
		)

		if err := rows.Scan(&id, &title, &desc, &price, &img, &size, &stock); err != nil {
			return nil, err
		}
		products = append(products, Products_admin{
			Id:    id,
			Image: img,
			Title: title,
			Price: price,
			Desc:  desc,
			Sizes: size,
			Stock: stock,
		})
	}

	return products, nil
}

// redux, dockerfile,
