package db

import (
	"context"
	"testing"
	"time"

	"github.com/Thanhbinh1905/go-training-bank/pkg/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T) *Transfer {
	fromAccount := createRandomAccount(t)
	toAccount := createRandomAccount(t)

	args := CreateTransferParams{
		FromAccountID: fromAccount.ID,
		ToAccountID:   toAccount.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, args.FromAccountID, transfer.FromAccountID)
	require.Equal(t, args.ToAccountID, transfer.ToAccountID)
	require.Equal(t, args.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreateAt)

	return &transfer
}

func TestCreateTransfer(t *testing.T) {
	createRandomTransfer(t)
}

func TestGetTransfer(t *testing.T) {
	entry1 := createRandomTransfer(t)
	entry2, err := testQueries.GetTransfer(context.Background(), entry1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.FromAccountID, entry2.FromAccountID)
	require.Equal(t, entry1.ToAccountID, entry2.ToAccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreateAt.Time, entry2.CreateAt.Time, time.Second)
}

func TestDeleteTransfer(t *testing.T) {
	entry1 := createRandomTransfer(t)
	err := testQueries.DeleteTransfer(context.Background(), entry1.ID)
	require.NoError(t, err)

	entry2, err := testQueries.GetTransfer(context.Background(), entry1.ID)
	require.Error(t, err)
	require.EqualError(t, err, "no rows in result set")
	require.Empty(t, entry2)
}
