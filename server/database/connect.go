package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DbConn *sql.DB

func ConnectDatabase(host, port, username, password, dbname string) error {
	dbSource := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8",
		username,
		password,
		host,
		port,
		dbname,
	)
	db, err := sql.Open("mysql", dbSource)
	if err != nil {
		panic(err)
	}
	DbConn = db
	return err
}
