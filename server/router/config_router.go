package router

import (
	"encoding/json"
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
	router.HandleFunc("/api/products", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}).Methods(http.MethodGet)
}

func routeUser(router *mux.Router) {
	router.HandleFunc("/login", ctrl.Login).Methods(http.MethodPost)
	router.HandleFunc("/register", ctrl.Register).Methods(http.MethodGet)
}
