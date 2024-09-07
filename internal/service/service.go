package service

import (
	"github.com/lks-go/pass-keeper/internal/lib/token"
	"github.com/lks-go/pass-keeper/internal/service/backend"
)

type ServerConfig struct {
	BinaryChunkSize int
}

type ServerDeps struct {
	Storage      backend.Storage
	PasswordHash backend.PasswordHash
	Token        *token.Token
	Crypt        backend.Crypt
}

func NewBackend(cfg ServerConfig, d ServerDeps) *backend.Service {
	return &backend.Service{
		BinaryChunkSize: cfg.BinaryChunkSize,

		Storage:  d.Storage,
		Password: d.PasswordHash,
		Token:    d.Token,
		Crypt:    d.Crypt,
	}
}
