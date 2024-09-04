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
	Crypt    Crypt
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

type Data struct {
	ID       int32
	Title    string
	Login    string
	Password string
}

func (s *Service) AddDataLoginPass(ctx context.Context, ownerLogin string, data Data) error {
	u, err := s.Storage.UserByLogin(ctx, ownerLogin)
	if err != nil {
		switch {
		case errors.Is(err, ErrNotFound):
			return ErrUserNotFound
		default:
			return fmt.Errorf("failed to get user by login")
		}
	}

	data.Login, err = s.Crypt.Encrypt(data.Login)
	if err != nil {
		return fmt.Errorf("failed to encrypt login: %w", err)
	}

	data.Password, err = s.Crypt.Encrypt(data.Password)
	if err != nil {
		return fmt.Errorf("failed to encrypt password: %w", err)
	}

	if err := s.Storage.AddLoginPass(ctx, u.ID, data); err != nil {
		return fmt.Errorf("failed to add data: %w", err)
	}

	return nil
}

func (s *Service) DataLoginPassList(ctx context.Context, ownerLogin string) ([]Data, error) {
	u, err := s.Storage.UserByLogin(ctx, ownerLogin)
	if err != nil {
		switch {
		case errors.Is(err, ErrNotFound):
			return nil, ErrUserNotFound
		default:
			return nil, fmt.Errorf("failed to get user by login")
		}
	}

	data, err := s.Storage.LoginPassList(ctx, u.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get data list: %w", err)
	}

	return data, nil
}
