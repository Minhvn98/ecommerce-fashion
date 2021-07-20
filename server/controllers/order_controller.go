package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	repo "github.com/Minhvn98/ecommerce-fashion/repository"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

func CreateNewOrder(w http.ResponseWriter, r *http.Request) {
	userId := context.Get(r, "id")
	id, _ := strconv.ParseInt(userId.(string), 10, 64)

	var orderInfo = make(map[string]string)
	err := json.NewDecoder(r.Body).Decode(&orderInfo)
	if nil != err {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "Error"})
		return
	}

	products := repo.GetProductsInCart(int(id))
	if len(products) > 0 {
		orderId := repo.CreateOrder(int(id), orderInfo, products)
		if orderInfo["payment"] == "Thanh toán momo" {
			url := PaymentWithMomo(int(id), int(orderId))
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]interface{}{"orderId": orderId, "url": url, "paymentType": 1})
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]interface{}{"orderId": orderId, "paymentType": 0})
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "Not have product in cart to checkout"})
	}

}

func GetOrderById(w http.ResponseWriter, r *http.Request) {
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

func GetOrdersByStatus(w http.ResponseWriter, r *http.Request) {
	userId := context.Get(r, "id")
	id, _ := strconv.ParseInt(userId.(string), 10, 64)

	vars := mux.Vars(r)
	status := vars["status"]

	orders := repo.GetOrdersByStatus(status, int(id))
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(orders)
}
