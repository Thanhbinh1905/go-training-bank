package service

import (
	"context"

	db "github.com/Thanhbinh1905/go-training-bank/internal/db/sqlc"
	"github.com/Thanhbinh1905/go-training-bank/internal/dto"
)

func (s *service) CreateEntry(ctx context.Context, req *dto.CreateTransferRequest) (*db.Transfer, error) {
	resp, err := s.store.CreateTransfer(ctx, db.CreateTransferParams{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
	})
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
