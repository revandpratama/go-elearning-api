package dto

type LoginRequest struct {
	UsernameOrEmail string `json:"username_or_Email" binding:"required" validate:"min=3,max=50"`
	Password        string `json:"password" binding:"required" validate:"min=6,max=20"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterRequest struct {
	Role            string `json:"role" binding:"required"`
	Name            string `json:"name" binding:"required" validate:"min=3,max=50"`
	Username        string `json:"username" binding:"required" validate:"min=3,max=50,alphanum"`
	Email           string `json:"email" binding:"required" validate:"email"`
	Phone           string `json:"phone" binding:"required" validate:"min=10,max=15,numeric"`
	Password        string `json:"password" binding:"required" validate:"min=8,max=50"`
	PasswordConfirm string `json:"password_confirm" binding:"required" validate:"min=8,max=50"`
}
