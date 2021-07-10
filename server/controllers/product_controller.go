package controller

import (
	"encoding/json"
	"net/http"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

func GetProductById(w http.ResponseWriter, r *http.Request) {

}
