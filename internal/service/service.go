package service

import (
	"context"

	db "github.com/Thanhbinh1905/go-training-bank/internal/db/sqlc"
	"github.com/Thanhbinh1905/go-training-bank/internal/dto"
)

type Service interface {
	CreateAccount(ctx context.Context, username string, req *dto.CreateAccountRequest) (*db.Account, error)
	GetAccount(ctx context.Context, req *dto.GetAccountRequest) (*db.Account, error)
	ListAccounts(ctx context.Context, owner string, req *dto.ListAccountsRequest) ([]db.Account, error)
	DeleteAccount(ctx context.Context, req *dto.DeleteAccountRequest) error

	CreateTransfer(ctx context.Context, req *dto.CreateTransferRequest) (*db.TransferTxResult, error)

	CreateUser(ctx context.Context, req *dto.CreateUserRequest) (*db.User, error)
	GetUser(ctx context.Context, req *dto.GetUserRequest) (*db.User, error)
}

type service struct {
	store db.Store
}

func NewService(store db.Store) Service {
	return &service{store: store}
}
