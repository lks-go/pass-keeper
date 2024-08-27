package service

import (
	"context"
	"fmt"
)

type hash interface {
	HashPassword(string) string
}

// NewUserRegistrarWithPassHash used on server to hash user password
func NewUserRegistrarWithPassHash(cfg *Config, s Storage) *UserRegistrar {
	return &UserRegistrar{
		cfg: &UserRegistrarConfig{
			UserPassSalt: cfg.UserPassSalt,
		},
		storage: s,
	}
}

type UserRegistrarWithPassHash struct {
	cfg     *UserRegistrarConfig
	storage Storage
	hash    hash
}

// Register registers a new user with his login and password
func (s *UserRegistrarWithPassHash) Register(ctx context.Context, u User) (string, error) {
	u.Password = s.hash.HashPassword(u.Password)

	if _, err := s.storage.Register(ctx, u); err != nil {
		return "", fmt.Errorf("filed to register user: %w", err)
	}

	return "", nil
}
