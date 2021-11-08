package user

import (
	//"mime/multipart"
)

type (
	// Request format
	SignUpReq struct {
		UserReq User
	}

	CheckUsernameReq struct {
		Username string `json:"username,omitempty"`
	}

	// Response format
	DefaultResponse struct {
		Status  bool   `json:"status"`
		Message string `json:"msg"`
	}
)
