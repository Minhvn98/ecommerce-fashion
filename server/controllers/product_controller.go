package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	repo "github.com/Minhvn98/ecommerce-fashion/repository"
	"github.com/gorilla/mux"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	limit, err1 := strconv.Atoi(r.FormValue("limit"))
	offset, err2 := strconv.Atoi(r.FormValue("offset"))
	if err1 != nil || err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "Invalid param"})
		return
	}

	products := repo.GetProducts(limit, offset)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idProduct, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "Invalid id"})
		return
	}
	product := repo.GetProductById(idProduct)
	if product.ID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "Product not exists"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func GetProductsByCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryId, err := strconv.Atoi(vars["id"])
	limit, err1 := strconv.Atoi(r.FormValue("limit"))
	offset, err2 := strconv.Atoi(r.FormValue("offset"))

	if err != nil || err1 != nil || err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "Invalid param"})
		return
	}

	if categoryId == 0 {
		products := repo.GetProducts(limit, offset)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(products)
		return
	}

	products := repo.GetProductsByCategory(categoryId, limit, offset)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

func GetProductsByTextSearch(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("text")
	limit, err1 := strconv.Atoi(r.FormValue("limit"))
	offset, err2 := strconv.Atoi(r.FormValue("offset"))

	if err1 != nil || err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "Invalid param"})
		return
	}

	products := repo.GetProductsByTextSearch(text, limit, offset)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}
