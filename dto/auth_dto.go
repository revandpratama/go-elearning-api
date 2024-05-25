package dto

type LoginRequest struct {
	UsernameOrEmail string `json:"username_or_Email" binding:"required" validate:"min=3,max=50"`
	Password        string `json:"password" binding:"required" validate:"min=6,max=20"`
}

type LoginResponse struct {
	Token string `json:"token"`
}






