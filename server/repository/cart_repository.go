package repository

import (
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

func GetProductsInCart(userId int) map[string]interface{} {
	results, err := db.DbConn.Query("SELECT id WHERE user_id = ?", userId)
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

	results2, err := db.DbConn.Query("SELECT id FROM cart_items WHERE cart_id = ?", cartId)
	if err != nil {
		fmt.Println(err)
	}

	var cart_id int
	defer results2.Close()
	for results2.Next() {
		err = results2.Scan(&cart_id)
		if err != nil {
			fmt.Println(err)
		}
	}

	return nil
}

func AddProductToCart(product map[string]int, userId int) {
	// Check user in cart
	results, err := db.DbConn.Query("SELECT id, user_id FROM carts WHERE user_id = ?", userId)
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

	if cartId == 0 {
		// Add new cart for user
		cartId = AddNewCartForUser(userId)
	}

	// insert, err := db.DbConn.Exec("INSERT INTO cart_items (cart_id, product_id, quantity) VALUES ( ?, ?, ? )", )
	// if err != nil {
	// 	fmt.Println(err)
	// }
}
