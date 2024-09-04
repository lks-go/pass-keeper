package server

import (
	"context"
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

type LoginPassData struct {
	ID       int32
	Title    string
	Login    string
	Password string
}

func (s *Service) AddDataLoginPass(ctx context.Context, ownerLogin string, data LoginPassData) (int32, error) {
	u, err := s.Storage.UserByLogin(ctx, ownerLogin)
	if err != nil {
		return 0, fmt.Errorf("failed to get user by login: %w", err)
	}

	data.Login, err = s.Crypt.Encrypt(data.Login)
	if err != nil {
		return 0, fmt.Errorf("failed to encrypt login: %w", err)
	}

	data.Password, err = s.Crypt.Encrypt(data.Password)
	if err != nil {
		return 0, fmt.Errorf("failed to encrypt password: %w", err)
	}

	id, err := s.Storage.AddLoginPass(ctx, u.ID, data)
	if err != nil {
		return 0, fmt.Errorf("failed to add data: %w", err)
	}

	return id, nil
}

func (s *Service) DataLoginPassList(ctx context.Context, ownerLogin string) ([]LoginPassData, error) {
	u, err := s.Storage.UserByLogin(ctx, ownerLogin)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by login: %w", err)
	}

	data, err := s.Storage.LoginPassList(ctx, u.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get data list: %w", err)
	}

	return data, nil
}

func (s *Service) DataLoginPass(ctx context.Context, ownerLogin string, ID int32) (*LoginPassData, error) {
	u, err := s.Storage.UserByLogin(ctx, ownerLogin)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by login: %w", err)
	}

	data, err := s.Storage.LoginPassByID(ctx, u.ID, ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get login&password: %w", err)
	}

	data.Login, err = s.Crypt.Decrypt(data.Login)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt login: %w", err)
	}

	data.Password, err = s.Crypt.Decrypt(data.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt password: %w", err)
	}

	return data, nil
}

type DataText struct {
	ID    int32
	Title string
	Text  string
}

func (s *Service) AddDataText(ctx context.Context, ownerLogin string, data DataText) (int32, error) {
	u, err := s.Storage.UserByLogin(ctx, ownerLogin)
	if err != nil {
		return 0, fmt.Errorf("failed to get user by login: %w", err)
	}

	data.Text, err = s.Crypt.Encrypt(data.Text)
	if err != nil {
		return 0, fmt.Errorf("failed to encrypt text: %w", err)
	}

	id, err := s.Storage.AddText(ctx, u.ID, data)
	if err != nil {
		return 0, fmt.Errorf("failed to add text to storage: %w", err)
	}

	return id, nil
}
