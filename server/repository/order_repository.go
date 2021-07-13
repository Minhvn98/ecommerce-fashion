package repository

import (
	"fmt"
	"strings"
	"time"

	db "github.com/Minhvn98/ecommerce-fashion/database"
	"github.com/Minhvn98/ecommerce-fashion/models"
)

func CreateOrder(userId int, orderInfo map[string]string, products []models.ProductInCart) {
	tx, err := db.DbConn.Begin()
	if err != nil {
		fmt.Println(err)
	}

	date := strings.Split(time.Now().String(), " ")[0]
	var orderId int64
	status := "Chờ thanh toán"
	totalPrice := 0
	// insert orders table
	{
		stmt, err := tx.Prepare(`INSERT INTO orders (user_id, status, payment, first_name, last_name, email, phone, location, total_price, created_at)
                     VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`)
		if err != nil {
			tx.Rollback()
			fmt.Println(err)
			return
		}
		defer stmt.Close()

		insert, err := stmt.Exec(userId, status, orderInfo["payment"], orderInfo["first_name"], orderInfo["last_name"], orderInfo["email"], orderInfo["phone"], orderInfo["location"], totalPrice, date)
		if err != nil {
			tx.Rollback()
			fmt.Println("insert orders table", err)
			return
		} else {
			orderId, err = insert.LastInsertId()
			if err != nil {
				tx.Rollback()
				fmt.Println(err)
				return
			}
		}
	}

	// insert for order_items table
	{
		stmt, err := tx.Prepare(`INSERT INTO order_items (product_id, order_id, quantity)
        VALUES(?, ?, ?);`)
		if err != nil {
			tx.Rollback()
			fmt.Println(err)
			return
		}
		defer stmt.Close()

		for _, product := range products {
			if product.Product.SalePercent != 0 {
				totalPrice += (product.Product.Price * int(product.Product.SalePercent) / 100) * product.Quantity
			} else {
				totalPrice += product.Product.Price * product.Quantity
			}
			_, err := stmt.Exec(product.Product.ID, orderId, product.Quantity)
			if err != nil {
				tx.Rollback()
				fmt.Println("insert for order_items table", err)
				return
			}
		}
	}
	//  update total price
	{
		_, err := tx.Exec("UPDATE orders SET total_price = ? WHERE id = ?", totalPrice, orderId)
		if err != nil {
			tx.Rollback()
			fmt.Println("update total price", err)
			return
		}
	}

	// Delete product in cart
	{
		for _, product := range products {
			_, err := tx.Exec("DELETE FROM cart_items WHERE product_id = ?", product.Product.ID)
			if err != nil {
				tx.Rollback()
				fmt.Println("Delete product in cart", err)
				return
			}
		}
	}

	tx.Commit()
}
