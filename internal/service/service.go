package service

import "github.com/lks-go/pass-keeper/internal/service/server"

type ServerDeps struct {
	Storage      server.Storage
	PasswordHash server.PasswordHash
}

func NewServer(d ServerDeps) *server.Service {
	return &server.Service{
		Storage:  d.Storage,
		Password: d.PasswordHash,
	}
}
