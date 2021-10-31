package user

import (
	"time"
	"strings"
)

type User struct {
	UserID    int    `json:"user_id,omitempty"`
	Username  string	`json:"username,omitempty"`
	Firstname string    `json:"firstname,omitempty"`
	Lastname  string    `json:"lastname,omitempty"`
	Password  string	`json:"password,omitempty"`
	Created_date time.Time `json:"created_date,omitempty"`
	Update_date  time.Time `json:"updated_date,omitempty"`
}

func splitName (fullname string) [] string{
	nameParts := strings.Split(fullname, " ")

	return nameParts
}