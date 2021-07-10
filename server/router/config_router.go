package router

import (
	"net/http"

	ctrl "github.com/Minhvn98/ecommerce-fashion/controllers"
	middle "github.com/Minhvn98/ecommerce-fashion/middlewares"
	"github.com/gorilla/mux"
)

func ConfigRouter() *mux.Router {
	router := mux.NewRouter()

	fs := http.FileServer(http.Dir("./public"))
	router.PathPrefix("/public").Handler(http.StripPrefix("/public", fs))
	router.Use(middle.CommonMiddleware)

	routerNoAuth := router.PathPrefix("/api/v1").Subrouter()
	routerAuth := router.PathPrefix("/api/v1").Subrouter()
	routerAuth.Use(middle.Authorization)

	RouterNoAuth(routerNoAuth)
	RouterAuth(routerAuth)

	return router
}

func RouterNoAuth(router *mux.Router) {
	// user
	router.HandleFunc("/login", ctrl.Login).Methods(http.MethodPost)
	router.HandleFunc("/register", ctrl.Register).Methods(http.MethodPost)

	// product
	router.HandleFunc("/products", ctrl.GetProduct).Methods(http.MethodGet)
	router.HandleFunc("/products/{id}", ctrl.GetProductById).Methods(http.MethodGet)
}

func RouterAuth(router *mux.Router) {
	// user
	router.HandleFunc("/logout", ctrl.Logout).Methods(http.MethodGet)

	// product
}
