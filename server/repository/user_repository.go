package repository

import (
	"fmt"
	"log"

	db "github.com/Minhvn98/ecommerce-fashion/database"
	"github.com/Minhvn98/ecommerce-fashion/models"
)

func GetUserById(id int) models.User {
	results, err := db.DbConn.Query("SELECT id, username, email, role FROM users WHERE id = ?", id)
	if err != nil {
		fmt.Println(err)
	}
	defer results.Close()

	var user models.User
	for results.Next() {
		err = results.Scan(&user.ID, &user.Name, &user.Email, &user.Role)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(user.ToString())
	}
	return user
}

func FindUserByUsername(username string) models.User {
	results, err := db.DbConn.Query("SELECT id, username, email, password, role FROM users WHERE username = ?", username)
	if err != nil {
		fmt.Println(err)
	}
	var user models.User
	defer results.Close()
	for results.Next() {
		err = results.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role)
		if err != nil {
			fmt.Println(err)
		}
	}
	return user
}

func CreateNewUser(user map[string]string) error {
	_, err := db.DbConn.Exec("INSERT INTO users (username, email, password, role ) VALUES (?, ?, ?, ?)", user["username"], user["email"], user["password"], user["role"])
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
