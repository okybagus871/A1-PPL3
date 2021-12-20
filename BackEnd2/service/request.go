package service

import (
	"time"
	"BackEnd/datastruct"
)

type (
	// Request format
	SignUpReq struct {
		UserReq datastruct.User
	}

	CheckUsernameReq struct {
		Username string `json:"username,omitempty"`
	}

	CheckEmailReq struct {
		Email string `json:"email,omitempty"`
	}

	ValidateEmailReq struct {
		Email	string `json:"email"`
		OTP		uint32 `json:"otp"`		
	}

	GetUserByEmailReq struct {
		Email	string `json:"email"`
	}

	// Response format
	DefaultResponse struct {
		Status  bool   `json:"status"`
		Message string `json:"msg"`
	}

	UserRes struct {
		Username		string		`json:"username,omitempty"`
		Password		string		`json:"password,omitempty"`
		Email			string		`json:"email,omitempty"`
		Name			string		`json:"name,omitempty"`
		Created_date	time.Time	`json:"created_date,omitempty"`
		Phonenumber     string    	`json:"phonenumber,omitempty"`
		Identity_type  	string    	`json:"identity_type,omitempty"`
		Identity_no    	string    	`json:"identity_no,omitempty"`
		Emergency_call 	string    	`json:"emergency_call,omitempty"`
		Address_ktp    	string    	`json:"address_ktp,omitempty"`
		Postal_code     string      `json:"postal_code,omitempty"`
	}
)