package router

import (
	"net/http"

	ctrl "github.com/Minhvn98/ecommerce-fashion/controller"
	"github.com/gorilla/mux"
)

func ConfigRouter() *mux.Router {
	router := mux.NewRouter()

	fs := http.FileServer(http.Dir("./public"))
	router.PathPrefix("/public").Handler(http.StripPrefix("/public", fs))

	routeProduct(router)
	routeUser(router)

	return router
}

func routeProduct(router *mux.Router) {
	router.HandleFunc("/api/products", ctrl.GetProduct).Methods(http.MethodGet)
}

func routeUser(router *mux.Router) {
	router.HandleFunc("/login", ctrl.Login).Methods(http.MethodPost)
	router.HandleFunc("/register", ctrl.Register).Methods(http.MethodPost)
}
