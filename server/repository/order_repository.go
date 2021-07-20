package repository

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	db "github.com/Minhvn98/ecommerce-fashion/database"
	"github.com/Minhvn98/ecommerce-fashion/models"
)

func CreateOrder(userId int, orderInfo map[string]string, products []models.ProductInCart) int64 {
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
			return 0
		}
		defer stmt.Close()

		insert, err := stmt.Exec(userId, status, orderInfo["payment"], orderInfo["first_name"], orderInfo["last_name"], orderInfo["email"], orderInfo["phone"], orderInfo["location"], totalPrice, date)
		if err != nil {
			tx.Rollback()
			fmt.Println("insert orders table", err)
			return 0
		} else {
			orderId, err = insert.LastInsertId()
			if err != nil {
				tx.Rollback()
				fmt.Println(err)
				return 0
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
			return 0
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
				return 0
			}
		}
	}
	//  update total price
	{
		_, err := tx.Exec("UPDATE orders SET total_price = ? WHERE id = ?", totalPrice, orderId)
		if err != nil {
			tx.Rollback()
			fmt.Println("update total price", err)
			return 0
		}
	}

	// Delete product in cart
	{
		for _, product := range products {
			_, err := tx.Exec("DELETE FROM cart_items WHERE product_id = ?", product.Product.ID)
			if err != nil {
				tx.Rollback()
				fmt.Println("Delete product in cart", err)
				return 0
			}
		}
	}

	tx.Commit()
	return orderId
}

func GetOrderById(id int) models.Order {
	var order models.Order
	{
		results, err := db.DbConn.Query("SELECT id, user_id, status, payment, first_name, last_name, email, phone, location, total_price, created_at FROM orders WHERE id = ? ", id)
		if err != nil {
			fmt.Println(err)
		}
		defer results.Close()
		for results.Next() {
			err = results.Scan(&order.Id, &order.UserId, &order.Status, &order.Payment, &order.FirstName, &order.LastName, &order.Email, &order.Phone, &order.Location, &order.TotalPrice, &order.CreatedAt)
			if err != nil {
				fmt.Println(err)
			}
		}
	}

	{
		results, err := db.DbConn.Query("SELECT product_id, quantity FROM order_items WHERE order_id = ? ", id)
		if err != nil {
			fmt.Println(err)
		}
		defer results.Close()
		products := make([]models.ProductInCart, 0)
		var productId, quantity int
		for results.Next() {
			err = results.Scan(&productId, &quantity)
			if err != nil {
				fmt.Println(err)
			}
			var product models.ProductInCart
			product.Product = GetProductById(productId)
			product.Quantity = quantity
			products = append(products, product)
		}
		order.Products = products
	}

	return order
}

func UpdateStatusOrder(status string, orderId int) {
	_, err := db.DbConn.Exec("UPDATE orders SET status = ? WHERE id = ?", status, orderId)
	if err != nil {
		fmt.Println("update status", err)
	}
}

func GetOrdersByStatus(status string, userId int) []map[string]interface{} {
	ordersId := make([]int, 0)
	orders := make([]map[string]interface{}, 0)
	{
		var results *sql.Rows
		var err error
		if status == "all" {
			results, err = db.DbConn.Query("SELECT id FROM orders WHERE user_id = ? ORDER BY id DESC ", userId)
		} else {
			results, err = db.DbConn.Query("SELECT id FROM orders WHERE user_id = ? AND status = ? ORDER BY id DESC", userId, status)
		}
		if err != nil {
			fmt.Println(err)
		}

		defer results.Close()
		id := 0
		for results.Next() {
			err = results.Scan(&id)
			ordersId = append(ordersId, id)
			if err != nil {
				fmt.Println(err)
			}
		}
	}

	{
		for _, id := range ordersId {
			results, err := db.DbConn.Query("SELECT product_id FROM order_items WHERE order_id = ?", id)
			if err != nil {
				fmt.Println(err)
			}
			defer results.Close()

			orderItem := make(map[string]interface{})
			var products []models.Product
			orderItem["orderId"] = id
			productId := 0
			for results.Next() {
				err = results.Scan(&productId)
				if err != nil {
					fmt.Println(err)
				}
				product := GetProductById(productId)
				products = append(products, product)
			}
			orderItem["products"] = products
			orders = append(orders, orderItem)
		}
	}
	return orders
}
