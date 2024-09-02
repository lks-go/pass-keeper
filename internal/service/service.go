package service

import (
	"github.com/lks-go/pass-keeper/internal/lib/token"
	"github.com/lks-go/pass-keeper/internal/service/server"
)

type ServerDeps struct {
	Storage      server.Storage
	PasswordHash server.PasswordHash
	Token        *token.Token
}

func NewServer(d ServerDeps) *server.Service {
	return &server.Service{
		Storage:  d.Storage,
		Password: d.PasswordHash,
		Token:    d.Token,
	}
}
