package models

import (
	"github.com/gin-gonic/gin"
)

type Order_Request struct {
	UserId          int            `json:"user_id" form:"user_id"`
	ShippingId      int            `json:"shipping_id" form:"shipping_id"`
	PaymentMethodId int            `json:"payment_method_id" form:"payment_method_id"`
	TotalOrder      uint64         `json:"total_order" form:"total_order"`
	NoOrder         string         `json:"no_order" form:"no_order"`
	StatusId        int            `json:"status_id" form:"status_id"`
	PromoId         int            `json:"promo_id" form:"promo_id"`
	Products        []ProductOrder `json:"products" form:"products"`
}

type ProductOrder struct {
	Id       int    `json:"id" form:"id"`
	SizeId   int    `json:"size_id" form:"size_id"`
	VarianId int    `json:"variant_id" form:"variant_id"`
	Qty      int    `json:"qty" form:"qty"`
	Total    uint64 `json:"total" form:"total"`
}

type AddCart struct {
	Product string `json:"product"`
	Status  string `json:"status"`
}

type Order struct{}

func (o *Order) CreateOrder(c *gin.Context, Uid int, req []ProductOrder) ([]AddCart, error) {

	var ress []AddCart
	for _, p := range req {
		var addCart AddCart

		var basePrice uint64
		Pg.QueryRow(c, `
			SELECT title, base_price
			FROM products
			WHERE id = $1
		`, p.Id).Scan(&addCart.Product, &basePrice)

		var sizePrice uint64
		Pg.QueryRow(c, `
			SELECT additional_price
			FROM sizes
			WHERE id = $1
		`, p.SizeId).Scan(&sizePrice)

		var variantPrice uint64
		Pg.QueryRow(c, `
			SELECT additional_price
			FROM variants
			WHERE id = $1
		`, p.VarianId).Scan(&variantPrice)

		subtot := basePrice + sizePrice + variantPrice

		_, err := Pg.Exec(c, `
			INSERT INTO carts(user_id, product_id, size_id, varian_id, qty, subtotal, product_name) VALUES
			($1, $2, $3, $$, $5, $6, $7);
		`, Uid, p.Id, p.SizeId, p.VarianId, p.Qty, subtot, addCart.Product)
		if err != nil {
			addCart.Status = "Gagal ditambahkan"
		}
		addCart.Status = "Berhasil ditambahkan"

		ress = append(ress, addCart)
	}

	return ress, nil
}
