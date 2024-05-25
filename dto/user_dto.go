package dto

type UserResponse struct {
	ID       int    `json:"id"`
	Role     string `json:"role"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}


type RegisterUserRequest struct {
	Name            string `json:"name" binding:"required" validate:"min=3,max=50"`
	Username        string `json:"username" binding:"required" validate:"min=3,max=50,alphanum"`
	Email           string `json:"email" binding:"required" validate:"email"`
	Phone           string `json:"phone" binding:"required" validate:"min=10,max=15,numeric"`
	Password        string `json:"password" binding:"required" validate:"min=8,max=50"`
	PasswordConfirm string `json:"password_confirm" binding:"required" validate:"min=8,max=50"`
}