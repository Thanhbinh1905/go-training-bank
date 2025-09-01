package token

import (
	"errors"
	"testing"
	"time"

	"github.com/Thanhbinh1905/go-training-bank/pkg/util"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"
)

func TestJWTMaker(t *testing.T) {
	maker, err := NewJWTMaker(util.RandomString(32))
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

	jwtPayload, ok := payload.(*JWTPayload)
	require.True(t, ok)

	require.NotZero(t, jwtPayload.ID)
	require.Equal(t, username, jwtPayload.Username)
	require.WithinDuration(t, issuedAt, jwtPayload.IssuedAt.Time, time.Second)
	require.WithinDuration(t, expiredAt, jwtPayload.ExpiresAt.Time, time.Second)
}

func TestExpiredJWTToken(t *testing.T) {
	maker, err := NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)

	username := util.RandomOwner()
	duration := -time.Minute // token expired

	token, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.True(t, errors.Is(err, jwt.ErrTokenExpired))
	require.Nil(t, payload)
}

func TestInvalidJWTTokenAlgNone(t *testing.T) {
	username := util.RandomOwner()
	duration := time.Minute

	payload, err := NewJWTPayload(username, duration)
	require.NoError(t, err)

	jwtPayload, ok := payload.(*JWTPayload)
	require.True(t, ok)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, jwtPayload)
	token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)

	maker, err := NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)

	payload, err = maker.VerifyToken(token)
	require.Error(t, err)
	require.True(t, errors.Is(err, ErrInvalidToken))
	require.Nil(t, payload)
}
