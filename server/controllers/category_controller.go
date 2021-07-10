package controllers

import (
	"encoding/json"
	"net/http"

	repo "github.com/Minhvn98/ecommerce-fashion/repository"
)

func GetCategories(w http.ResponseWriter, r *http.Request) {
	categories := repo.GetCategories()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(categories)
}
