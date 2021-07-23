package models

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	ID                int               `json:"id"`
	Name              string            `json:"name"`
	Description       string            `json:"description"`
	Price             int               `json:"price"`
	SalePercent       float32           `json:"sale_percent"`
	Quantity          int               `json:"quantity"`
	Category          Category          `json:"category"`
	ProductImages     []ProductImage    `json:"product_images"`
	ProductProperties []ProductProperty `json:"product_properties"`
}

func (product *Product) String() string {
	if product == nil {
		return ""
	}

	b, err := json.Marshal(product)
	if err != nil {
		fmt.Println("Cannot convert to json: ", err)
		return ""
	}
	
	return string(b)
}
