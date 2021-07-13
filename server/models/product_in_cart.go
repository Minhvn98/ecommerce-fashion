package models

type ProductInCart struct {
	Product  Product `json:"product_info"`
	Quantity int     `json:"quantity_in_cart"`
}
