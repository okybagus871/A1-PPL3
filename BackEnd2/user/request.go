package user

import (
	//"mime/multipart"
)

type (
	// Request format
	SignUpReq struct {
		// Username string `json:"username" binding:"required"`
		// Password string `json:"password" binding:"required"`
		// Name 	 string `json:"name" binding:"required"`
		User userProfile
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
