package user

type RegisterUserRequest struct {
	FirstName   string `json:"first_name" binding:"required"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Address     string `json:"address" binding:"required"`
	Pin         string `json:"pin" binding:"required"`
}

type LoginUserRequest struct {
	PhoneNumber string `json:"phone_number" binding:"required"`
	Pin         string `json:"pin" binding:"required"`
}

type User struct {
	UserID   string `json:"user_id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type RegisterUserResponse struct {
	UserID      string `json:"user_id"`
	FirstName   string `json:"first_name" `
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	CreatedAt   string `json:"created_at"`
}

type LoginUserResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type UpdateUserRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name"`
	Address   string `json:"address" binding:"required"`
	UserID    string
}

type UpdateUserResponse struct {
	UserID    string `json:"user_id"`
	FirstName string `json:"first_name" `
	LastName  string `json:"last_name"`
	Address   string `json:"address"`
	UpdatedAt string `json:"updated_at"`
}
