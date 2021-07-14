package models

type Order struct {
	Id         int             `json:"id"`
	UserId     int             `json:"user_id"`
	Status     string          `json:"status"`
	FirstName  string          `json:"first_name"`
	LastName   string          `json:"last_name"`
	Email      string          `json:"email"`
	Phone      string          `json:"phone"`
	Location   string          `json:"location"`
	TotalPrice int             `json:"total_price"`
	Payment    string          `json:"payment"`
	Products   []ProductInCart `json:"product"`
	CreatedAt  string          `json:"created_at"`
	UpdatedAt  string          `json:"updated_at"`
}
