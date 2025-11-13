package models

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type Transaction_Request struct {
	UserId           int    `json:"-"`
	Name             string `json:"name"`
	Address          string `json:"adress"`
	Phone            string `json:"phone"`
	Email            string `json:"email"`
	PaymentMethod_id int    `json:"payment_method_id"`
	Delivery_id      int    `json:"delivery_id"`
	Invoice          string `json:"invoice"`
	Total            uint64 `json:"total"`
}

type Carts struct {
	User_id    int
	Product_id int
	Size_id    int
	Varian_Id  int
	Qty        int
	SubTotal   float64
	Name       string
}

type Transactions struct{}

func (t Transactions) CreateTransactions(c context.Context, req Transaction_Request) (err error) {
	conn, err := Pg.Acquire(c)
	if err != nil {
		return fmt.Errorf("failed to acquire connection: %w", err)
	}
	defer conn.Release()

	tx, err := conn.Begin(c)
	if err != nil {
		return fmt.Errorf("failed starting transaction: %w", err)
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback(c)
		}
	}()

	// generate invoice
	invoice := fmt.Sprintf("VIA-%d-%d", time.Now().Unix(), req.UserId)
	fmt.Println("Invoice:", invoice)

	// ambil data dari cart
	rows, err := tx.Query(c, `
		SELECT product_id, size_id, varian_id, qty, subtotal, name
		FROM carts
		WHERE user_id = $1`, req.UserId)
	if err != nil {
		return fmt.Errorf("failed to query carts: %w", err)
	}
	defer rows.Close()

	var subTotal float64

	for rows.Next() {
		var prodId, sizeId, varianId, qty int
		var subtotal float64
		var name string

		if err = rows.Scan(&prodId, &sizeId, &varianId, &qty, &subtotal, &name); err != nil {
			return fmt.Errorf("error scanning cart: %w", err)
		}

		// insert ke orders_products
		_, err = tx.Exec(c, `
			INSERT INTO orders_products (invoice, user_id, product_id, size_id, varian_id, qty, subtotal, name)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		`, invoice, req.UserId, prodId, sizeId, varianId, qty, subtotal, name)
		if err != nil {
			return fmt.Errorf("failed inserting order_products: %w", err)
		}

		// update stok produk
		_, err = tx.Exec(c, `
			UPDATE products SET stock = stock - $1 WHERE id = $2
		`, qty, prodId)
		if err != nil {
			return fmt.Errorf("failed updating stock: %w", err)
		}

		subTotal += subtotal
	}

	// insert ke orders
	_, err = tx.Exec(c, `
		INSERT INTO orders (user_id, email, fullname, phone, address, payment_method_id, delivery_id, total_order, invoice, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	`, req.UserId, req.Email, req.Name, req.Phone, req.Address, req.PaymentMethod_id, req.Delivery_id, subTotal, invoice, time.Now(), time.Now())
	if err != nil {
		return fmt.Errorf("failed inserting orders: %w", err)
	}

	// hapus cart
	_, _ = tx.Exec(c, `DELETE FROM carts WHERE user_id = $1`, req.UserId)

	if err = tx.Commit(c); err != nil {
		return fmt.Errorf("failed committing transaction: %w", err)
	}

	return nil
}

type History_res struct {
	Id      int    `json:"id"`
	Invoice string `json:"invoice"`
	Date    string `json:"date"`
	Total   uint64 `json:"total"`
	Status  string `json:"status"`
	Image   string `json:"image"`
}

type History_req struct {
	User_id int `json:"-"`
	Month   int `json:"month"`
	Status  int `json:"status"`
	Page    int `json:"date"`
	Limit   int `json:"total"`
}

func (t Transactions) GetHistory(c context.Context, req History_req) ([]History_res, error) {
	rows, err := Pg.Query(c, `
		SELECT 
			o.invoice, 	
			o.created_at, 	
			o.total_order,
			s.name,
			MIN(pi.image)
		FROM orders o
		LEFT JOIN orders_products op ON op.invoice = o.invoice
		LEFT JOIN products_images pi ON pi.product_id = op.product_id
		LEFT JOIN status s ON s.id = o.status_id
		WHERE o.user_id = $1
		GROUP BY o.invoice, o.created_at, o.total_order, s.name
		ORDER BY o.created_at DESC
	`, req.User_id)

	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	var ress []History_res
	id := 0
	for rows.Next() {
		var (
			invoice string
			created time.Time
			total   uint64
			status  string
			image   sql.NullString
		)

		rows.Scan(&invoice, &created, &total, &status, &image)

		ress = append(ress, History_res{
			Id:      id,
			Invoice: invoice,
			Date:    created.Format("2006-01-02"),
			Total:   total,
			Status:  status,
			Image:   image.String,
		})
		id++
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}
	return ress, nil
}
