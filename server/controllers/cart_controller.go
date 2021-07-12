package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Minhvn98/ecommerce-fashion/middlewares"
	repo "github.com/Minhvn98/ecommerce-fashion/repository"
)

func AddProductToCart(w http.ResponseWriter, r *http.Request) {
	userId := middlewares.GetUserIdFromToken(w, r, "customer")
	if userId == 0 {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "Access is not allowed"})
		return
	}
	var product = make(map[string]int)
	err := json.NewDecoder(r.Body).Decode(&product)
	if nil != err {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "Must use int type to send"})
		return
	}

	repo.AddProductToCart(product, userId)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func GetProductsInCart(w http.ResponseWriter, r *http.Request) {
	userId := middlewares.GetUserIdFromToken(w, r, "customer")
	if userId == 0 {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "Access is not allowed"})
		return
	}

	products := repo.GetProductsInCart(userId)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}
