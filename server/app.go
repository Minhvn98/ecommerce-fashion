package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Minhvn98/ecommerce-fashion/config"
	"github.com/Minhvn98/ecommerce-fashion/database"
	"github.com/Minhvn98/ecommerce-fashion/router"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Load config
	err := config.LoadConfig()
	if err != nil {
		panic(err.Error())
	}

	// Connect database
	dbName := config.Config.Db.Database
	dbUserName := config.Config.Db.User
	dbPass := config.Config.Db.Pass
	dbHost := config.Config.Db.Host
	dbPort := config.Config.Db.Port

	database.ConnectDatabase(dbHost, dbPort, dbUserName, dbPass, dbName)
	database.InitAdminAccount()

	// Disconnect database
	defer database.DbConn.Close()

	// Config service
	srv := &http.Server{
		Handler:      router.ConfigRouter(),
		Addr:         "127.0.0.1:3000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Server will start at http://localhost:3000/")
	log.Fatal(srv.ListenAndServe())

}
