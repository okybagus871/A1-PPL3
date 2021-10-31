package user

type UserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Firstname string `json:"firstname" binding:"required"`
}
