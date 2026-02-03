package dto

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	UserId  int    `json:"user_id"`
	Token   string `json:"token"`
	Message string `json:"message"`
}
