package repository

import (
	"errors"
	"fmt"

	db "github.com/Minhvn98/ecommerce-fashion/database"
)

func AddNewCartForUser(userId int) int {
	insert, err := db.DbConn.Exec("INSERT INTO carts (user_id) VALUES ( ? )", userId)
	if err != nil {
		fmt.Println(err)
	}

	cartId, err := insert.LastInsertId()
	if err != nil {
		fmt.Println(err)
	}
	return int(cartId)
}

func GetCartIdByUserId(userId int) int {
	results, err := db.DbConn.Query("SELECT id FROM carts WHERE user_id = ?", userId)
	if err != nil {
		fmt.Println(err)
	}
	var cartId int
	defer results.Close()
	for results.Next() {
		err = results.Scan(&cartId)
		if err != nil {
			fmt.Println(err)
		}
	}
	return cartId
}

func GetProductsInCart(userId int) []map[string]interface{} {
	cartId := GetCartIdByUserId(userId)

	results2, err := db.DbConn.Query("SELECT product_id, quantity FROM cart_items WHERE cart_id = ?", cartId)
	if err != nil {
		fmt.Println(err)
	}
	defer results2.Close()

	products := make([]map[string]interface{}, 0)
	var productId, quantity int
	for results2.Next() {
		err = results2.Scan(&productId, &quantity)
		if err != nil {
			fmt.Println(err)
		}
		product := make(map[string]interface{})
		product["product_info"] = GetProductById(productId)
		product["quantity_in_cart"] = quantity
		products = append(products, product)
	}

	return products
}

func IsExistsProductInCart(productId uint, cartId int) bool {
	results, err := db.DbConn.Query("SELECT id FROM cart_items WHERE cart_id = ? and product_id = ?", cartId, productId)
	if err != nil {
		fmt.Println(err)
	}
	defer results.Close()
	var id int
	for results.Next() {
		err = results.Scan(&id)
		if err != nil {
			fmt.Println(err)
		}
	}
	return id != 0
}

func UpdateCart(product map[string]uint, userId int) error {
	var cartId int
	cartId = GetCartIdByUserId(userId)

	if cartId == 0 {
		// Add new cart for user
		cartId = AddNewCartForUser(userId)
	}

	check := IsExistsProductInCart(product["product_id"], cartId)
	if !check {
		// Add new product in cert
		if product["quantity"] == 0 {
			return errors.New("message: quantity must different 0")
		}
		_, err := db.DbConn.Exec("INSERT INTO cart_items (cart_id, product_id, quantity) VALUES ( ?, ?, ? )", cartId, product["product_id"], product["quantity"])
		if err != nil {
			fmt.Println(err)
			return err
		}
	} else {
		if product["quantity"] == 0 {
			// Delete product when quantity equal 0
			_, err := db.DbConn.Exec("DELETE FROM cart_items WHERE cart_id = ? and product_id = ?", cartId, product["product_id"])
			if err != nil {
				fmt.Println(err)
				return err
			}
		} else {
			// Update product in cart
			_, err := db.DbConn.Exec("UPDATE cart_items SET quantity = ? WHERE cart_id = ? and product_id = ?", product["quantity"], cartId, product["product_id"])
			if err != nil {
				fmt.Println(err)
				return err
			}
		}
	}
	return nil
}
