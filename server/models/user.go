package models

import (
	"fmt"
	"strconv"
)

type User struct {
	ID     int    `json:"id"`
	Name   string `json:"username"`
	Mobile string `json:"image"`
	Role   string `json:"role"`
}

func (user *User) ToString() string {
	return fmt.Sprintf("Id: %s, Name: %s, Mobile: %s, Role: %s", strconv.Itoa(user.ID), user.Name, user.Mobile, user.Role)

}
