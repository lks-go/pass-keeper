package service

import (
	"github.com/lks-go/pass-keeper/internal/lib/token"
	"github.com/lks-go/pass-keeper/internal/service/backend"
)

type ServerDeps struct {
	Storage      backend.Storage
	PasswordHash backend.PasswordHash
	Token        *token.Token
	Crypt        backend.Crypt
}

func NewBackend(d ServerDeps) *backend.Service {
	return &backend.Service{
		Storage:  d.Storage,
		Password: d.PasswordHash,
		Token:    d.Token,
		Crypt:    d.Crypt,
	}
}
