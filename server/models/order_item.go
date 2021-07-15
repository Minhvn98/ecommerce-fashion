package models

type OrderItem struct {
	Id        int `json:"id"`
	ProductId int `json:"product_id"`
	OrderId   int `json:"order_id"`
	Quantity  int `json:"quantity"`
}
