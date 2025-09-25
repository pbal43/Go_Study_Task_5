package auth

import (
	"fmt"
	"time"
	"toDoList/internal/server/auth/auth_errors"

	"github.com/golang-jwt/jwt/v5"
)

type ParseOptions struct {
	ExpectedIssuer   string
	ExpectedAudience string
	AllowMethods     []string
	Leeway           time.Duration
}

func (hs HS256Signer) ParseAccessToken(token string, opt ParseOptions) (*Claims, error) {
	claims := Claims{}
	tok, err := jwt.ParseWithClaims(
		token,
		&claims,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return hs.Secret, nil
		},
		jwt.WithIssuer(opt.ExpectedIssuer),
		jwt.WithLeeway(opt.Leeway),
		jwt.WithAudience(opt.ExpectedAudience),
		jwt.WithValidMethods(opt.AllowMethods),
	)
	if err != nil {
		return nil, err
	}
	if !tok.Valid {
		return nil, auth_errors.ErrorInvalidAccessToken
	}
	return &claims, nil
}

func (hs HS256Signer) ParseRefreshToken(token string, opt ParseOptions) (*jwt.RegisteredClaims, error) {
	claims := jwt.RegisteredClaims{}
	tok, err := jwt.ParseWithClaims(
		token,
		&claims,
		func(token *jwt.Token) (interface{}, error) {
			return hs.Secret, nil
		},
		jwt.WithIssuer(opt.ExpectedIssuer),
		jwt.WithLeeway(opt.Leeway),
		jwt.WithAudience(opt.ExpectedAudience),
		jwt.WithValidMethods(opt.AllowMethods),
	)
	if err != nil {
		return nil, err
	}
	if !tok.Valid {
		return nil, auth_errors.ErrorInvalidRefreshToken
	}
	return &claims, nil
}
