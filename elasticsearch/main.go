package main

import (
	"log"
	"net/http"

	"es-server/routes"
)

func main() {

	log.Println("Server will start at http://localhost:4000")

	log.Fatal(http.ListenAndServe(":4000", routes.ConfigRouter()))
}
