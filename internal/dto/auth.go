package dto

// RegisterRequest struct for register user
type RegisterRequest struct {
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=6"`
	FirstName string `json:"first_name" binding:"required,min=2,max=32"`
	LastName  string `json:"last_name" binding:"required,min=2,max=32"`
	Phone     string `json:"phone"`
}

// LoginRequest struct for login user
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// RefreshTokenRequest struct for refresh token
type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

// AuthResponse struct for return response after auth
type AuthResponse struct {
	User         UserResponse `json:"user"`
	AccessToken  string       `json:"access_token"`
	RefreshToken string       `json:"refresh_token"`
}

// UserResponse struct for return fields that belong to user
type UserResponse struct {
	ID        uint   `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Role      string `json:"role"`
	IsActive  bool   `json:"is_active"`
}

// UpdateProfileRequest struct for profile update
type UpdateProfileRequest struct {
	FirstName string `json:"first_name" binding:"required,min=2,max=32"`
	LastName  string `json:"last_name" binding:"required,min=2,max=32"`
	Phone     string `json:"phone"`
}
