package token

import (
	"errors"
	"testing"
	"time"

	"github.com/Thanhbinh1905/go-training-bank/pkg/util"
	"github.com/stretchr/testify/require"
)

func TestPasetoMaker(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	username := util.RandomOwner()
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	pasetoPayload, ok := payload.(*PasetoPayload)
	require.True(t, ok)

	require.NotZero(t, pasetoPayload.ID)
	require.Equal(t, username, pasetoPayload.Username)
	require.WithinDuration(t, issuedAt, pasetoPayload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, pasetoPayload.ExpiredAt, time.Second)
}

func TestExpiredPasetoToken(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	username := util.RandomOwner()
	duration := -time.Minute // token expired

	token, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.True(t, errors.Is(err, ErrExpiredToken))
	require.Nil(t, payload)
}
