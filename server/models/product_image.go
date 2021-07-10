package models

type ProductImage struct {
	ID        int    `json:"id"`
	ProductId int    `json:"product_id"`
	Uri       string `json:"uri"`
}
