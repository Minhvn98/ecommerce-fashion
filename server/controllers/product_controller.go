package controllers

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Minhvn98/ecommerce-fashion/models"
	repo "github.com/Minhvn98/ecommerce-fashion/repository"
	"github.com/Minhvn98/ecommerce-fashion/services/email"
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

func UploadFile(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	record, _ := reader.ReadAll()
	fmt.Println(len(record[0]))
	for _, line := range record {
		var product models.Product
		product.Name = line[0]
		product.Description = line[1]
		product.Price, _ = strconv.Atoi(line[2])
		s, _ := strconv.ParseFloat(line[3], 32)
		product.SalePercent = float32(s)
		categoryId, _ := strconv.Atoi(line[4])
		product.Quantity, _ = strconv.Atoi(line[5])
		repo.AddNewProduct(product, categoryId)
	}
	email.SendMail("levankhanh.dev@gmail.com", models.Mail{Subject: "Thông báo về update sản phẩm", Body: "Bạn đã tải lên sản phẩm thành công"})

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "Tải lên sản phẩm thành công"})
}

// func UploadFile(w http.ResponseWriter, r *http.Request) {

// 	common.Publish(rabbitmq.Chanel, "UploadProduct", "dev")

// 	msgs := common.Consume(rabbitmq.Chanel, "UploadProduct")
// 	go func() {
// 		for d := range msgs {
// 			myString := string(d.Body[:])
// 			fmt.Println(myString)
// 		}
// 	}()

// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode("ok")
// }
