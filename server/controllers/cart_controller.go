package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	repo "github.com/Minhvn98/ecommerce-fashion/repository"
	"github.com/gorilla/context"
)

func UpdateCart(w http.ResponseWriter, r *http.Request) {
	userId := context.Get(r, "id")
	id, _ := strconv.ParseInt(userId.(string), 10, 64)
	var product = make(map[string]uint)
	err := json.NewDecoder(r.Body).Decode(&product)
	if nil != err {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "Must use uint type to send"})
		return
	}

	e := repo.UpdateCart(product, int(id))
	if e != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "Error when update, please try again"})
		return
	}

	products := repo.GetProductsInCart(int(id))
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

func GetCart(w http.ResponseWriter, r *http.Request) {
	userId := context.Get(r, "id")
	id, _ := strconv.ParseInt(userId.(string), 10, 64)
	products := repo.GetProductsInCart(int(id))

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}
