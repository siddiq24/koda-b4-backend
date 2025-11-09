package models

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
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

type Order struct {
	Pg *pgxpool.Pool
}

func (o *Order) CreateOrder(c *gin.Context, req Order_Request) error {

	fmt.Println(req)
	QueryOrder := fmt.Sprintf("INSERT INTO orders( user_id, shipping_id, no_order, status_id, promo_id, total_order) VALUES(%d,%d,'%s',%d,%d, 0) RETURNING id;",
		req.UserId, req.ShippingId, req.NoOrder, req.StatusId, req.PromoId)
	var orderId int
	if err := o.Pg.QueryRow(c, QueryOrder).Scan(&orderId); err != nil {
		return err
	}

	var totalOrder uint64
	for _, product := range req.Products {
		if _, err := o.Pg.Exec(c, `INSERT INTO orders_products(order_id, product_id, size_id, varian_id, qty) VALUES ($1, $2, $3, $4, $5)`, orderId, product.Id, product.SizeId, product.VarianId, product.Qty); err != nil {
			return err
		}
		totalOrder += product.Total
	}

	if _, err := o.Pg.Exec(c, `UPDATE orders SET total_order = $1 WHERE id = $2`, totalOrder, orderId); err != nil {
		return err
	}
	return nil
}
