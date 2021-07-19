package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Minhvn98/ecommerce-fashion/middlewares"
	repo "github.com/Minhvn98/ecommerce-fashion/repository"
	"github.com/gorilla/mux"
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
	if len(products) > 0 {
		orderId := repo.CreateOrder(userId, orderInfo, products)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{"orderId": orderId})

		PaymentWithMomo(userId, int(orderId))
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "Not have product in cart"})
	}

}

func GetOrderById(w http.ResponseWriter, r *http.Request) {
	userId := middlewares.GetUserIdFromToken(w, r, "customer")
	if userId == 0 {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "Access is not allowed"})
		return
	}

	vars := mux.Vars(r)
	orderId, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "Invalid id"})
		return
	}

	order := repo.GetOrderById(orderId)
	if order.Id == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "Order not exists"})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(order)
	}

}

func UpdateStatusOrder(w http.ResponseWriter, r *http.Request) {
	userId := middlewares.GetUserIdFromToken(w, r, "admin")
	if userId == 0 {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "Access is not allowed"})
		return
	}

	vars := mux.Vars(r)
	orderId, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "Invalid id"})
		return
	}

	var userReq = make(map[string]string)
	err = json.NewDecoder(r.Body).Decode(&userReq)
	if nil != err {
		fmt.Println(err)
	}

	repo.UpdateStatusOrder(userReq["status"], orderId)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "Update status sucessfully"})
}
