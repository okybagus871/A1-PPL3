package user

import (
	"time"
)

type userProfile struct {
		UserID         uint32    `json:"user_id,omitempty"`
		Username       string    `json:"username,omitempty"`
		Email          string    `json:"email,omitempty"`
		Name		   string	 `json:"name,omitempty"`
		// Firstname      string    `json:"firstname,omitempty"`
		// Lastname       string    `json:"lastname,omitempty"`
		Phonenumber    string    `json:"phonenumber,omitempty"`
		Password       string    `json:"password,omitempty"`
		Created_date   time.Time `json:"created_date,omitempty"`
		Updated_date   time.Time `json:"updated_date,omitempty"`
		Email_verified bool      `json:"email_verified,omitempty"`
		Image_file     string    `json:"image_file,omitempty"`
		Identity_type  string    `json:"identity_type,omitempty"`
		Identity_no    string    `json:"identity_no,omitempty"`
		Emergency_call string    `json:"emergency_call,omitempty"`
		Address_ktp    string    `json:"address_ktp,omitempty"`
		Domisili       string    `json:"domisili,omitempty"`
}