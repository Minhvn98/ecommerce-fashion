package repository

import (
	"fmt"
	"log"

	db "github.com/Minhvn98/ecommerce-fashion/database"
	"github.com/Minhvn98/ecommerce-fashion/models"
)

func GetUser() {
	results, err := db.DbConn.Query("SELECT id, username, email ,role FROM users")
	if err != nil {
		panic(err.Error())
	}
	defer results.Close()
	for results.Next() {
		var user models.User
		err = results.Scan(&user.ID, &user.Name, &user.Email, &user.Role)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(user.ToString())
	}
}

func FindUserByUsername(username string) models.User {
	results, err := db.DbConn.Query("SELECT id, username, email, password, role FROM users WHERE username = ?", username)
	if err != nil {
		panic(err.Error())
	}
	var user models.User
	defer results.Close()
	for results.Next() {
		err = results.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role)
		if err != nil {
			panic(err.Error())
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
