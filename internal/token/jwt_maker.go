package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var minSecretKeySize = 32

type JWTMaker struct {
	secretKey string
}

func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minSecretKeySize)
	}
	return &JWTMaker{secretKey: secretKey}, nil
}

func (maker *JWTMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewJWTPayload(username, duration)
	if err != nil {
		return "", err
	}

	jwtPayload, ok := payload.(*JWTPayload) // ép kiểu từ interface -> struct
	if !ok {
		return "", ErrInvalidTokenPayload
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtPayload)
	return token.SignedString([]byte(maker.secretKey))
}

func (maker *JWTMaker) VerifyToken(token string) (Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		// Kiểm tra token ký bằng HS256
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return []byte(maker.secretKey), nil
	}

	parsedToken, err := jwt.ParseWithClaims(token, &JWTPayload{}, keyFunc)
	if err != nil {
		return nil, err
	}

	jwtpayload, ok := parsedToken.Claims.(*JWTPayload)
	if !ok {
		return nil, ErrInvalidTokenPayload
	}

	return jwtpayload, nil
}
