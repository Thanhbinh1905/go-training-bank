package service

import (
	"context"
	"fmt"

	db "github.com/Thanhbinh1905/go-training-bank/internal/db/sqlc"
	"github.com/Thanhbinh1905/go-training-bank/internal/dto"
)

func (s *service) CreateAccount(ctx context.Context, username string, req *dto.CreateAccountRequest) (*db.Account, error) {
	resp, err := s.store.CreateAccount(ctx, db.CreateAccountParams{
		Owner:    username,
		Balance:  0,
		Currency: req.Currency,
	})
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (s *service) GetAccount(ctx context.Context, req *dto.GetAccountRequest) (*db.Account, error) {
	account, err := s.store.GetAccount(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (s *service) ListAccounts(ctx context.Context, owner string, req *dto.ListAccountsRequest) ([]db.Account, error) {

	accounts, err := s.store.ListAccounts(ctx, db.ListAccountsParams{
		Owner:  owner,
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	})
	if err != nil {
		return nil, fmt.Errorf("service.ListAccounts: %w", err)
	}

	return accounts, nil
}

func (s *service) DeleteAccount(ctx context.Context, req *dto.DeleteAccountRequest) error {
	return s.store.DeleteAccount(ctx, req.ID)
}
