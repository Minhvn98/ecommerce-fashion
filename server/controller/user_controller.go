package controller

import (
	"encoding/json"
	"net/http"
	// repo "github.com/Minhvn98/ecommerce-fashion/repository"
)

func Login(w http.ResponseWriter, r *http.Request) {
	// repo.GetUser()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

func Register(w http.ResponseWriter, r *http.Request) {

}
