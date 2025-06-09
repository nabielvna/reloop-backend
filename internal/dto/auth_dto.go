package dto

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	UserName string `json:"userName"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}

type UserResponse struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	UserName string `json:"userName"`
	Role     string `json:"role"`
}
