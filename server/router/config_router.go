package router

import (
	"net/http"

	ctrl "github.com/Minhvn98/ecommerce-fashion/controllers"
	middle "github.com/Minhvn98/ecommerce-fashion/middlewares"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func ConfigRouter() http.Handler {
	router := mux.NewRouter()

	fs := http.FileServer(http.Dir("./public"))
	router.PathPrefix("/public").Handler(http.StripPrefix("/public", fs))
	// router.Use(middle.CommonMiddleware)

	routerNoAuth := router.PathPrefix("/api/v1").Subrouter()
	routerAuth := router.PathPrefix("/api/v1").Subrouter()
	routerAuth.Use(middle.Authentication)

	RouterNoAuth(routerNoAuth)
	RouterAuth(routerAuth)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)

	return handler
}

func RouterNoAuth(router *mux.Router) {
	// user
	router.HandleFunc("/login", ctrl.Login).Methods(http.MethodPost)
	router.HandleFunc("/register", ctrl.Register).Methods(http.MethodPost)

	// product
	router.HandleFunc("/products", ctrl.GetProducts).Queries("limit", "{limit:[0-9]+}", "offset", "{offset:[0-9]+}").Methods(http.MethodGet)
	router.HandleFunc("/products/{id:[0-9]+}", ctrl.GetProductById).Methods(http.MethodGet)
	router.HandleFunc("/products/search", ctrl.GetProductsByTextSearch).Queries("text", "{text:.*}", "limit", "{limit:[0-9]+}", "offset", "{offset:[0-9]+}").Methods(http.MethodGet)

	// category
	router.HandleFunc("/categories", ctrl.GetCategories).Methods(http.MethodGet)
	router.HandleFunc("/categories/{id}/products", ctrl.GetProductsByCategory).Queries("limit", "{limit:[0-9]+}", "offset", "{offset:[0-9]+}").Methods(http.MethodGet)
}

func RouterAuth(router *mux.Router) {
	// user
	router.HandleFunc("/user/token", ctrl.GetUserByToken).Methods(http.MethodGet)
	router.HandleFunc("/logout", ctrl.Logout).Methods(http.MethodGet)

	// cart
	router.HandleFunc("/cart", ctrl.GetCart).Methods(http.MethodGet)
	router.HandleFunc("/cart", ctrl.UpdateCart).Methods(http.MethodPut)

	// Checkout
	router.HandleFunc("/checkout", ctrl.CreateNewOrder).Methods(http.MethodPost)
}
