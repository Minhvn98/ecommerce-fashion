package models

type ProductProperty struct {
	ID        int    `json:"id"`
	ProductId int    `json:"product_id"`
	Key       string `json:"key"`
	Value     string `json:"value"`
}
