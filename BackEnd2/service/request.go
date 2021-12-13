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
	}
)