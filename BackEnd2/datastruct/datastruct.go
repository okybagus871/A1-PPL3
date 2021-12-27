package datastruct

import (
	"time"
)

type User struct {
	UserID         uint32    `json:"user_id,omitempty"`
	Username       string    `json:"username,omitempty"`
	Email          string    `json:"email,omitempty"`
	Name           string    `json:"name,omitempty"`
	Token_hash     string    `json:"token_hash,omitempty"`
	OTP            int       `json:"otp,omitempty"`
	Phonenumber    string    `json:"phonenumber,omitempty"`
	Password       string    `json:"password,omitempty"`
	Created_date   string	 `json:"created_date,omitempty"`
	Updated_date   time.Time `json:"updated_date,omitempty"`
	Email_verified bool      `json:"email_verified,omitempty"`
	Image_file     string    `json:"image_file,omitempty"`
	Identity_type  string    `json:"identity_type,omitempty"`
	Identity_no    string    `json:"identity_no,omitempty"`
	Emergency_call string    `json:"emergency_call,omitempty"`
	Address_ktp    string    `json:"address_ktp,omitempty"`
	Postal_code    string    `json:"postal_code,omitempty"`
	Domisili       string    `json:"domisili,omitempty"`
}
