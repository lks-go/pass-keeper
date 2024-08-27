package service

import (
	"context"
	"fmt"
)

type UserRegistrarConfig struct {
	UserPassSalt string
}

type Storage interface {
	Register(ctx context.Context, u User) (string, error)
}

func NewUserRegistrar(cfg *Config, s Storage) *UserRegistrar {
	return &UserRegistrar{
		cfg: &UserRegistrarConfig{
			UserPassSalt: cfg.UserPassSalt,
		},
		storage: s,
	}
}

type UserRegistrar struct {
	cfg     *UserRegistrarConfig
	storage Storage
}

// TODO идея заключается в том что бы один и тот же сервис использовать для сервера и клиента
// клиент через интерфейс будет дергать сервер, а сервер через интерфейс будет дергать БД

// Register registers a new user with his login and password
func (s *UserRegistrar) Register(ctx context.Context, u User) (string, error) {
	if _, err := s.storage.Register(ctx, u); err != nil {
		return "", fmt.Errorf("filed to register user: %w", err)
	}

	return "", nil
}
