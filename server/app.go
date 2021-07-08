package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Minhvn98/ecommerce-fashion/config"
	"github.com/Minhvn98/ecommerce-fashion/database"
	"github.com/Minhvn98/ecommerce-fashion/router"
)

func main() {
	// Load config
	err := config.LoadConfig()
	if err != nil {
		panic(err.Error())
	}

	// Connect database
	dbName := "crawler_web"
	dbUserName := "root"
	dbPass := "123456"
	dbHost := "127.0.0.1"
	dbPort := "3306"

	database.ConnectDatabase(dbHost, dbPort, dbUserName, dbPass, dbName)

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
