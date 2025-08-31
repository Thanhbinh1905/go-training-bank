package service

import (
	"context"

	db "github.com/Thanhbinh1905/go-training-bank/internal/db/sqlc"
	"github.com/Thanhbinh1905/go-training-bank/internal/dto"
)

type Service interface {
	CreateAccount(ctx context.Context, req *dto.CreateAccountRequest) (*db.Account, error)
}

type service struct {
	store *db.Store
}

func NewService(store *db.Store) Service {
	return &service{store: store}
}
