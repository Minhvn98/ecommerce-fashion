package main

import (
	"log"
	"net/http"

	"github.com/Minhvn98/ecommerce-fashion/config"
	"github.com/Minhvn98/ecommerce-fashion/database"
	"github.com/Minhvn98/ecommerce-fashion/router"
	"github.com/Minhvn98/ecommerce-fashion/services/rabbitmq"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Load config
	err := config.LoadConfig()
	if err != nil {
		panic(err.Error())
	}

	// Run RabbitMQ
	go func() {
		rabbitmq.RunRabbitMQ()
	}()

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

	// Run service
	Addr := config.Config.Server.Host + ":" + config.Config.Server.Post
	log.Println("Server will start at http://" + Addr + "/")

	log.Fatal(http.ListenAndServe(":3000", router.ConfigRouter()))

}
