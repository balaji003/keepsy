package auth

import "time"

type UserCredentials struct {
	UserID       int       `json:"user_id"`
	PasswordHash string    `json:"-"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone,omitempty"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Identifier string `json:"identifier"` // Email or Phone
	Password   string `json:"password"`
}
