package service

import (
	"time"
	"BackEnd/datastruct"
)

type (
	// Request format
	UserProfileReq struct {
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

	UpdatePasswordReq struct {
		Email	string `json:"email"`
		Password string `json:"password"`
	}

	// Response format
	DefaultResponse struct {
		Status  bool   `json:"status"`
		Message string `json:"msg"`
	}

	UserRes struct {
		Username		string		`json:"username"`
		Password		string		`json:"password"`
		Email			string		`json:"email"`
		Name			string		`json:"name"`
		Created_date	time.Time	`json:"created_date"`
		Phonenumber     string    	`json:"phonenumber"`
		Identity_type  	string    	`json:"identity_type"`
		Identity_no    	string    	`json:"identity_no"`
		Emergency_call 	string    	`json:"emergency_call"`
		Address_ktp    	string    	`json:"address_ktp"`
		Postal_code     string      `json:"postal_code"`
	}

	UserPassRes struct {
		Password		string		`json:"password,omitempty"`
	}
)