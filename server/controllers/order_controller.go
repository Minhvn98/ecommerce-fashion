package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Minhvn98/ecommerce-fashion/middlewares"
	repo "github.com/Minhvn98/ecommerce-fashion/repository"
)

func CreateNewOrder(w http.ResponseWriter, r *http.Request) {
	userId := middlewares.GetUserIdFromToken(w, r, "customer")
	if userId == 0 {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "Access is not allowed"})
		return
	}

	var orderInfo = make(map[string]string)
	err := json.NewDecoder(r.Body).Decode(&orderInfo)
	if nil != err {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "Error"})
		return
	}

	products := repo.GetProductsInCart(userId)

	repo.CreateOrder(userId, orderInfo, products)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(orderInfo)

}
