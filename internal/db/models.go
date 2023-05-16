package db

import "time"

type User struct {
	User_Id  string    `json:"id" bson:"User_Id,omitempty"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}

type SignupRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}
