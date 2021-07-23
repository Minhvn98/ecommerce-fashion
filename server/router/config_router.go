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

	routerNoAuth := router.PathPrefix("/api/v1").Subrouter()
	routerAuth := router.PathPrefix("/api/v1").Subrouter()
	routerAuth.Use(middle.Authentication)

	RouterNoAuth(routerNoAuth)
	RouterAuth(routerAuth)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:*", "http://192.168.8.235:*", "http://118.69.210.244:*", "http://192.168.0.6:*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATH", "DELETE"},
	})

	handler := c.Handler(router)

	return handler
}

func RouterNoAuth(router *mux.Router) {
	// user
	router.Path("/login").Methods(http.MethodPost).HandlerFunc(ctrl.Login)
	router.Path("/register").Methods(http.MethodPost).HandlerFunc(ctrl.Register)

	// product
	router.Path("/products").Queries("limit", "{limit:[0-9]+}", "offset", "{offset:[0-9]+}").Methods(http.MethodGet).HandlerFunc(ctrl.GetProducts)
	router.Path("/products/{id:[0-9]+}").Methods(http.MethodGet).HandlerFunc(ctrl.GetProductById)
	router.Path("/products/search").Queries("text", "{text:.*}", "limit", "{limit:[0-9]+}", "offset", "{offset:[0-9]+}").Methods(http.MethodGet).HandlerFunc(ctrl.GetProductsByTextSearch)

	// category
	router.Path("/categories").Methods(http.MethodGet).HandlerFunc(ctrl.GetCategories)
	router.Path("/categories/{id}/products").Queries("limit", "{limit:[0-9]+}", "offset", "{offset:[0-9]+}").Methods(http.MethodGet).HandlerFunc(ctrl.GetProductsByCategory)

	// Payment
	router.Path("/payment").Methods(http.MethodGet).HandlerFunc(ctrl.GetPayment)
	router.Path("/payment").Methods(http.MethodPost).HandlerFunc(ctrl.PostPayment)
}

func RouterAuth(router *mux.Router) {
	// user
	router.Path("/user/token").Methods(http.MethodGet).HandlerFunc(ctrl.GetUserByToken)
	router.Path("/logout").Methods(http.MethodGet).HandlerFunc(ctrl.Logout)

	// product
	router.Path("/admin/product/upload").Methods(http.MethodPost).HandlerFunc(middle.Authorization(ctrl.UploadFile, "admin"))

	// cart
	router.Path("/cart").Methods(http.MethodGet).HandlerFunc(middle.Authorization(ctrl.GetCart, "customer"))
	router.Path("/cart").Methods(http.MethodPut).HandlerFunc(middle.Authorization(ctrl.UpdateCart, "customer"))

	// Order
	router.Path("/order-detail/{id:[0-9]+}").Methods(http.MethodGet).HandlerFunc(middle.Authorization(ctrl.GetOrderById, "customer", "admin"))
	router.Path("/order/{id:[0-9]+}").Methods(http.MethodPut).HandlerFunc(middle.Authorization(ctrl.UpdateStatusOrder, "admin"))
	router.Path("/checkout").Methods(http.MethodPost).HandlerFunc(middle.Authorization(ctrl.CreateNewOrder, "customer"))
	router.Path("/order").Queries("status", "{status:.*}").Methods(http.MethodGet).HandlerFunc(middle.Authorization(ctrl.GetOrdersByStatus, "customer"))
	router.Path("/orders").Methods(http.MethodGet).HandlerFunc(middle.Authorization(ctrl.GetAllOrder, "admin"))
}
