package models

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
