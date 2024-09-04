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

type DataLoginPass struct {
	ID       int32
	Title    string
	Login    string
	Password string
}

func (s *Service) AddDataLoginPass(ctx context.Context, ownerLogin string, data *DataLoginPass) (int32, error) {
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

func (s *Service) DataLoginPassList(ctx context.Context, ownerLogin string) ([]DataLoginPass, error) {
	u, err := s.Storage.UserByLogin(ctx, ownerLogin)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by login: %w", err)
	}

	data, err := s.Storage.LoginPassList(ctx, u.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get login & pass list: %w", err)
	}

	return data, nil
}

func (s *Service) DataLoginPass(ctx context.Context, ownerLogin string, ID int32) (*DataLoginPass, error) {
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

func (s *Service) AddDataText(ctx context.Context, ownerLogin string, data *DataText) (int32, error) {
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

func (s *Service) DataTextList(ctx context.Context, ownerLogin string) ([]DataText, error) {
	u, err := s.Storage.UserByLogin(ctx, ownerLogin)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by login: %w", err)
	}

	data, err := s.Storage.TextList(ctx, u.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get text list: %w", err)
	}

	return data, nil
}

func (s *Service) DataText(ctx context.Context, ownerLogin string, ID int32) (*DataText, error) {
	u, err := s.Storage.UserByLogin(ctx, ownerLogin)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by login: %w", err)
	}

	data, err := s.Storage.TextByID(ctx, u.ID, ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get text: %w", err)
	}

	data.Text, err = s.Crypt.Decrypt(data.Text)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt text: %w", err)
	}

	return data, nil
}

type DataCard struct {
	ID      int32
	Title   string
	Number  string
	Owner   string
	ExpDate string
	CVCCode string
}

func (s *Service) AddDataCard(ctx context.Context, ownerLogin string, data *DataCard) (int32, error) {
	u, err := s.Storage.UserByLogin(ctx, ownerLogin)
	if err != nil {
		return 0, fmt.Errorf("failed to get user by login: %w", err)
	}

	if err := s.encryptCardData(data); err != nil {
		return 0, fmt.Errorf("failed to encrypt card data: %w", err)
	}

	id, err := s.Storage.AddCard(ctx, u.ID, data)
	if err != nil {
		return 0, fmt.Errorf("failed to add card data: %w", err)
	}

	return id, nil
}

func (s *Service) encryptCardData(data *DataCard) error {
	var err error

	data.Number, err = s.Crypt.Encrypt(data.Number)
	if err != nil {
		return fmt.Errorf("failed to encrypt card number: %w", err)
	}

	data.Owner, err = s.Crypt.Encrypt(data.Owner)
	if err != nil {
		return fmt.Errorf("failed to encrypt card owner: %w", err)
	}

	data.ExpDate, err = s.Crypt.Encrypt(data.ExpDate)
	if err != nil {
		return fmt.Errorf("failed to encrypt expiration date: %w", err)
	}

	data.CVCCode, err = s.Crypt.Encrypt(data.CVCCode)
	if err != nil {
		return fmt.Errorf("failed to encrypt cvc code: %w", err)
	}

	return nil
}

func (s *Service) DataCardList(ctx context.Context, ownerLogin string) ([]DataCard, error) {
	u, err := s.Storage.UserByLogin(ctx, ownerLogin)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by login: %w", err)
	}

	data, err := s.Storage.CardList(ctx, u.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get card list: %w", err)
	}

	return data, nil
}
