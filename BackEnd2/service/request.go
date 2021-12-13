package service

import (
	//"time"
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

	// Response format
	DefaultResponse struct {
		Status  bool   `json:"status"`
		Message string `json:"msg"`
	}
)