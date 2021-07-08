package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Minhvn98/ecommerce-fashion/utils"
	_ "github.com/go-sql-driver/mysql"
)

var DbConn *sql.DB

func ConnectDatabase(host, port, username, password, dbname string) {
	dbSource := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		username,
		password,
		host,
		port,
		dbname,
	)

	db, err := sql.Open("mysql", dbSource)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	DbConn = db
}

func InitAdminAccount() {
	results, err := DbConn.Query("SELECT username FROM users WHERE username = 'admin'")
	if err != nil {
		panic(err.Error())
	}
	var username string
	defer results.Close()
	for results.Next() {

		err = results.Scan(&username)
		if err != nil {
			panic(err.Error())
		}
	}

	if username == "" {
		username = "admin"
		password := utils.HashPassword("admin")
		role := "admin"

		_, err = DbConn.Exec(`INSERT INTO users (username, password, email, role) VALUES ( ?, ?, ?, ?)`, username, password, "", role)
		if err != nil {
			log.Fatal(err)
		}
	}
}
