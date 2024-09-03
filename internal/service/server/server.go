package server

import (
	"context"
	"errors"
	"fmt"

	"github.com/lks-go/pass-keeper/internal/lib/token"
)

type Service struct {
	Storage  Storage
	Password PasswordHash
	Token    *token.Token
}

// RegisterUser registers a new user with his login and password
func (s *Service) RegisterUser(ctx context.Context, login string, password string) (string, error) {
	userId, err := s.Storage.RegisterUser(ctx, login, s.Password.Hash(password))
	if err != nil {
		return "", fmt.Errorf("failed to add user to storage: %w", err)
	}

	return userId, nil
}

func (s *Service) AuthUser(ctx context.Context, login string, password string) (string, error) {
	u, err := s.Storage.UserByLogin(ctx, login)
	if err != nil {
		return "", fmt.Errorf("failed to get user login: %w", err)
	}

	if s.Password.Hash(password) != u.PasswordHash {
		return "", ErrUsersPasswordNotMatch
	}

	token, err := s.Token.BuildNewJWTToken(login)
	if err != nil {
		return "", fmt.Errorf("failed to build token: %w", err)
	}

	return token, nil
}

type DataLoginPass struct {
	Title    string
	Login    string
	Password string
}

func (s *Service) AddDataLoginPass(ctx context.Context, ownerLogin string, data DataLoginPass) error {
	u, err := s.Storage.UserByLogin(ctx, ownerLogin)
	if err != nil {
		switch {
		case errors.Is(err, ErrNotFound):
			return ErrUserNotFound
		default:
			return fmt.Errorf("failed to get user by login")
		}
	}

	// TODO encrypt data
	if err := s.Storage.AddDataLoginPass(ctx, u.ID, data); err != nil {
		return fmt.Errorf("failed to add data: %w", err)
	}

	return nil
}
