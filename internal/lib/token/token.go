package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Token errors
var (
	ErrInvalidToken   = errors.New("invalid token")
	ErrTokenExpired   = errors.New("token expired")
	ErrEmptySecretKey = errors.New("secret key must not be empty")
)

// Claims embeds jwt.RegisteredClaims
type Claims struct {
	jwt.RegisteredClaims
	Login string
}

const (
	defaultTokenExp = time.Hour * 60
)

func New(sk string, exp time.Duration) (*Token, error) {
	if sk == "" {
		return nil, ErrEmptySecretKey
	}

	if exp == 0 {
		exp = defaultTokenExp
	}

	return &Token{secretKey: sk, exp: exp}, nil
}

type Token struct {
	secretKey string
	exp       time.Duration
}

// BuildNewJWTToken builds new jwt to userID
func (t *Token) BuildNewJWTToken(login string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(t.exp)),
		},
		Login: login,
	})

	tokenString, err := token.SignedString([]byte(t.secretKey))
	if err != nil {
		return "", fmt.Errorf("failed to get signed string: %w", err)
	}

	return tokenString, nil
}

// ParseJWTToken validates jwt
func (t *Token) ParseJWTToken(token string) (*Claims, error) {
	claims := Claims{}
	parsedToken, err := jwt.ParseWithClaims(token, &claims, func(jt *jwt.Token) (interface{}, error) {
		if jt.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("unexpected signing method: %v", jt.Header["alg"])
		}

		return []byte(t.secretKey), nil
	})
	if err != nil && !errors.Is(err, jwt.ErrTokenExpired) {
		return nil, fmt.Errorf("failed to parse jwt: %w", err)
	}

	if errors.Is(err, jwt.ErrTokenExpired) {
		return nil, ErrTokenExpired
	}

	if !parsedToken.Valid {
		return nil, ErrInvalidToken
	}

	return &claims, nil
}
