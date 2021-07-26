package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	ctrl "es-server/controllers"
)

func ConfigRouter() http.Handler {
	router := mux.NewRouter()

	searchRoute := router.PathPrefix("/api/v1").Subrouter()

	SearchRouter(searchRoute)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:*", "http://192.168.8.235:*", "http://118.69.210.244:*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATH", "DELETE"},
	})

	handler := c.Handler(router)

	return handler
}

func SearchRouter(router *mux.Router) {

	router.Path("/products/search").Queries("q", "{q:.*}", "limit", "{limit:[0-9]+}", "offset", "{offset:[0-9]+}").Methods(http.MethodGet).HandlerFunc(ctrl.SearchProduct)

	router.Path("/products/search").Queries("q", "{q:.*}").Methods(http.MethodGet).HandlerFunc(ctrl.SearchProduct)

	// index product
	router.Path("/products").Methods(http.MethodPost).HandlerFunc(ctrl.AddProduct)

	router.Path("/products/{id:[0-9]+}").Methods(http.MethodPut).HandlerFunc(ctrl.UpdateProduct)

	router.Path("/products/{id:[0-9]+}").Methods(http.MethodDelete).HandlerFunc(ctrl.DeleteProduct)

	router.Path("/products/add-products").Methods(http.MethodPost).HandlerFunc(ctrl.AddProducts)

}
