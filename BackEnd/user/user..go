package user

import (
	"time"
)

type User struct {
	Id    int    	`json:"user_id,omitempty"` // UserID
	Username  string	`json:"username,omitempty"`
	Name	  string    `json:"name,omitempty"`
	Password  string	`json:"password,omitempty"`
	Created_date time.Time `json:"created_date,omitempty"`
	Update_date  time.Time `json:"updated_date,omitempty"`
}