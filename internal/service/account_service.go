package service

import (
	"context"

	db "github.com/Thanhbinh1905/go-training-bank/internal/db/sqlc"
	"github.com/Thanhbinh1905/go-training-bank/internal/dto"
)

func (s *service) CreateAccount(ctx context.Context, req *dto.CreateAccountRequest) (*db.Account, error) {
	resp, err := s.store.CreateAccount(ctx, db.CreateAccountParams{
		Owner:    req.Owner,
		Balance:  0,
		Currency: req.Currency,
	})
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
