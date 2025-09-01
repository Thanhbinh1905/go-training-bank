package dto

import (
	"time"

	db "github.com/Thanhbinh1905/go-training-bank/internal/db/sqlc"
)

type CreateUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type GetUserRequest struct {
	Username string `uri:"username" binding:"required,alphanum"`
}

type UserResponse struct {
	Username string    `json:"username"`
	FullName string    `json:"full_name"`
	Email    string    `json:"email"`
	CreateAt time.Time `json:"create_at"`
}

type LoginUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginUserResponse struct {
	AccessToken string       `json:"access_token"`
	User        UserResponse `json:"user"`
}

func NewUserResponse(user *db.User) *UserResponse {
	return &UserResponse{
		Username: user.Username,
		FullName: user.FullName,
		Email:    user.Email,
		CreateAt: user.CreateAt,
	}
}
