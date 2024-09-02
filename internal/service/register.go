package service

import (
	"context"
	"fmt"
)

type UserRegistrarConfig struct {
	UserPassSalt string
}

type userStorage interface {
	AddUser(ctx context.Context, login string, passwordHash string) (string, error)
}

type passwordHasher interface {
	Hash(pass string) string
}

func NewUserRegistrar(cfg *Config, s userStorage, ph passwordHasher) *UserRegistrar {
	return &UserRegistrar{
		cfg: &UserRegistrarConfig{
			UserPassSalt: cfg.UserPassSalt,
		},
		userStorage: s,
		password:    ph,
	}
}

type UserRegistrar struct {
	cfg         *UserRegistrarConfig
	userStorage userStorage
	password    passwordHasher
}

// TODO идея заключается в том что бы один и тот же сервис использовать для сервера и клиента
// клиент через интерфейс будет дергать сервер, а сервер через интерфейс будет дергать БД

// RegisterUser registers a new user with his login and password
func (s *UserRegistrar) RegisterUser(ctx context.Context, u User) (string, error) {
	userId, err := s.userStorage.AddUser(ctx, u.Login, s.password.Hash(u.Password))
	if err != nil {
		return "", fmt.Errorf("failed to add user to storage: %w", err)
	}

	return userId, nil
}
