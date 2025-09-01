package service

import (
	"context"

	db "github.com/Thanhbinh1905/go-training-bank/internal/db/sqlc"
	"github.com/Thanhbinh1905/go-training-bank/internal/dto"
)

func (s *service) CreateUser(ctx context.Context, req *dto.CreateUserRequest) (*db.User, error) {
	resp, err := s.store.CreateUser(ctx, db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: req.Password,
		FullName:       req.FullName,
		Email:          req.Email,
	})
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (s *service) GetUser(ctx context.Context, req *dto.GetUserRequest) (*db.User, error) {
	user, err := s.store.GetUser(ctx, req.Username)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
