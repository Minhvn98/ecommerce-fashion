package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	repo "github.com/Minhvn98/ecommerce-fashion/repository"
	"github.com/Minhvn98/ecommerce-fashion/utils"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var userReq = make(map[string]string)
	err := json.NewDecoder(r.Body).Decode(&userReq)
	if nil != err {
		fmt.Println(err)
	}

	user := repo.FindUserByUsername(userReq["username"])

	// Check username
	if user.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "Username not found"})
		return
	}

	// Check password
	isMatchPassword := utils.CheckPasswordHash(userReq["password"], user.Password)
	if !isMatchPassword {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "Incorrect password"})
		return
	}

	// Create token
	token, err := utils.CreateToken(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Create cookie
	utils.SetCookieHandler(w, r, "token", token)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":       user.ID,
		"username": user.Name,
		"email":    user.Email,
		"role":     user.Role})
}

func Register(w http.ResponseWriter, r *http.Request) {
	var userReq = make(map[string]string)
	err := json.NewDecoder(r.Body).Decode(&userReq)
	if nil != err {
		log.Fatal(err)
	}
	//// validated ////
	if userReq["username"] == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "Username cannot be empty"})
		return
	}
	if userReq["password"] == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "Password cannot be empty"})
		return
	}
	if userReq["comfirm_password"] != userReq["password"] {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "Password not match"})
		return
	}
	user := repo.FindUserByUsername(userReq["username"])
	if user.ID != 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "Username already exists"})
		return
	}

	userReq["password"] = utils.HashPassword(userReq["password"])
	userReq["role"] = "customer"

	err = repo.CreateNewUser(userReq)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "Error"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "Register sucessfully, please login to shoping"})
}

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:     "token",
		Value:    "",
		Path:     "/",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "Logout sucessfully"})
}
