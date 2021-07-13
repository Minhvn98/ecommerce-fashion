package models

type Order struct {
	Id         int    `json:"id"`
	User_id    int    `json:"user_id"`
	Status     string `json:"status"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Location   string `json:"location"`
	TotalPrice int    `json:"total_price"`
}
