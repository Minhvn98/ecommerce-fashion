package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"es-server/models"
	"es-server/pkg"
	repo "es-server/repository"

	"github.com/gorilla/mux"
)

const (
	indexName = "ecommerce-fashion"
)

var url = "http://localhost:9200"
var esclient, _ = pkg.NewESClient(url)
var pm = repo.NewProductManager(esclient)

func init() {
	// Use the IndexExists service to check if a specified index exists.
	exists, err := esclient.IndexExists(indexName).Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}
	if !exists {
		// Create a new index.
		mapping := `
{
	"settings":{
		"number_of_shards":1,
		"number_of_replicas":0
	}
	
}
`
		createIndex, err := esclient.CreateIndex(indexName).Body(mapping).Do(context.Background())
		if err != nil {
			// Handle error
			panic(err)
		}
		if !createIndex.Acknowledged {
			// Not acknowledged
			fmt.Println("hehe")
		}
	}
}

func SearchProduct(w http.ResponseWriter, r *http.Request) {
	query := r.FormValue("q")

	resultSearchProducts := pm.SearchProducts(query)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resultSearchProducts)
}

func AddProduct(w http.ResponseWriter, r *http.Request) {
	// var userReq = make(map[string]interface{})
	var product *models.Product

	err := json.NewDecoder(r.Body).Decode(&product)
	if nil != err {
		fmt.Println(err)
	}

	err = pm.AddProduct(product)
	if err != nil {
		fmt.Println("Cannot insert product: ", product, err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func AddProducts(w http.ResponseWriter, r *http.Request) {
	// var userReq = make(map[string]interface{})
	var products *[]models.Product

	err := json.NewDecoder(r.Body).Decode(&products)
	if nil != err {
		fmt.Println(err)
	}

	insertedProductCount := 0
	for _, p := range *products {
		err = pm.AddProduct(&p)

		if err != nil {
			fmt.Println("Cannot insert book: ", p, err)
			continue
		}
		insertedProductCount++
	}

	msg := "Insert success " + strconv.Itoa(insertedProductCount) + " products"
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(msg)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var product *models.Product

	err := json.NewDecoder(r.Body).Decode(&product)
	if nil != err {
		fmt.Println(err)
	}

	err = pm.UpdateProduct(product)
	if err != nil {
		fmt.Println("Cannot update product: ", product, err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {

	var params = mux.Vars(r)
	idProduct := params["id"]

	err := pm.DeleteProduct(idProduct)
	if err != nil {
		fmt.Println("Cannot delete product: ", idProduct, err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Delete product success")
}
