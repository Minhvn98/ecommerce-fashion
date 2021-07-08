package repository

import (
	"fmt"

	db "github.com/Minhvn98/ecommerce-fashion/database"
	"github.com/Minhvn98/ecommerce-fashion/models"
)

func GetUser() {
	results, err := db.DbConn.Query("SELECT * FROM movies")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var user models.User
		err = results.Scan(&user.ID, &user.Name, &user.Mobile, &user.Role)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(user.ToString())
	}
}
