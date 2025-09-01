package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	pasword := RandomString(6)

	hashedPassword1, err := Hash(pasword)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword1)

	err = CheckHash(pasword, hashedPassword1)
	require.NoError(t, err)

	wrongPassword := RandomString(6)
	err = CheckHash(wrongPassword, hashedPassword1)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	hashedPassword2, err := Hash(pasword)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword2)
	require.NotEqual(t, hashedPassword1, hashedPassword2)
}
