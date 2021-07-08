package models

import (
	"fmt"
	"strconv"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func (user *User) ToString() string {
	return fmt.Sprintf("Id: %s, Name: %s, Email: %s, Role: %s", strconv.Itoa(user.ID), user.Name, user.Email, user.Role)

}
