package backend

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"

	"github.com/lks-go/pass-keeper/internal/lib/token"
	"github.com/lks-go/pass-keeper/internal/service/entity"
)

type Service struct {
	BinaryChunkSize int

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
		return "", entity.ErrUsersPasswordNotMatch
	}

	token, err := s.Token.BuildNewJWTToken(login)
	if err != nil {
		return "", fmt.Errorf("failed to build token: %w", err)
	}

	return token, nil
}

func (s *Service) AddDataLoginPass(ctx context.Context, ownerLogin string, data *entity.DataLoginPass) (int32, error) {
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

func (s *Service) DataLoginPassList(ctx context.Context, ownerLogin string) ([]entity.DataLoginPass, error) {
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

func (s *Service) DataLoginPass(ctx context.Context, ownerLogin string, ID int32) (*entity.DataLoginPass, error) {
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

func (s *Service) AddDataText(ctx context.Context, ownerLogin string, data *entity.DataText) (int32, error) {
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

func (s *Service) DataTextList(ctx context.Context, ownerLogin string) ([]entity.DataText, error) {
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

func (s *Service) DataText(ctx context.Context, ownerLogin string, ID int32) (*entity.DataText, error) {
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

func (s *Service) AddDataCard(ctx context.Context, ownerLogin string, data *entity.DataCard) (int32, error) {
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

func (s *Service) DataCardList(ctx context.Context, ownerLogin string) ([]entity.DataCard, error) {
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

func (s *Service) DataCard(ctx context.Context, ownerLogin string, ID int32) (*entity.DataCard, error) {
	u, err := s.Storage.UserByLogin(ctx, ownerLogin)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by login: %w", err)
	}

	data, err := s.Storage.CardByID(ctx, u.ID, ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get card: %w", err)
	}

	if err := s.decryptCardData(data); err != nil {
		return nil, fmt.Errorf("failed to decrypt card data: %w", err)
	}

	return data, nil
}

func (s *Service) encryptCardData(data *entity.DataCard) error {
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

func (s *Service) decryptCardData(data *entity.DataCard) error {
	var err error

	data.Number, err = s.Crypt.Decrypt(data.Number)
	if err != nil {
		return fmt.Errorf("failed to decrypt card number: %w", err)
	}

	data.Owner, err = s.Crypt.Decrypt(data.Owner)
	if err != nil {
		return fmt.Errorf("failed to decrypt card owner: %w", err)
	}

	data.ExpDate, err = s.Crypt.Decrypt(data.ExpDate)
	if err != nil {
		return fmt.Errorf("failed to decrypt expiration date: %w", err)
	}

	data.CVCCode, err = s.Crypt.Decrypt(data.CVCCode)
	if err != nil {
		return fmt.Errorf("failed to decrypt cvc code: %w", err)
	}

	return nil
}

func (s *Service) AddDataBinary(ctx context.Context, ownerLogin string, binary *entity.DataBinary) (int32, error) {
	u, err := s.Storage.UserByLogin(ctx, ownerLogin)
	if err != nil {
		return 0, fmt.Errorf("failed to get user by login: %w", err)
	}

	g, ctx := errgroup.WithContext(ctx)
	cnt := 0
	chunkNum := 0
	chunk := make([]byte, 0, s.BinaryChunkSize)

	binID, err := s.Storage.AddBinary(ctx, u.ID, binary)
	if err != nil {
		return 0, fmt.Errorf("failed to add binary: %w", err)
	}

	for b := range binary.Body {
		select {
		case <-ctx.Done():
			return 0, ctx.Err()
		default:
		}

		chunk = append(chunk, b)
		cnt++

		if cnt == s.BinaryChunkSize {
			cnt = 0
			chunkNum++

			func(orderNumber int) {
				chunkToEncrypt := string(chunk)
				chunk = chunk[:0]

				g.Go(func() error {
					encryptedChunk, err := s.Crypt.Encrypt(chunkToEncrypt)
					if err != nil {
						return fmt.Errorf("failed to encrypt chunk %d: %w", orderNumber, err)
					}

					err = s.Storage.AddBinaryChunk(ctx, binID, encryptedChunk, orderNumber)
					if err != nil {
						return fmt.Errorf("failed to add binary chunk %d: %w", orderNumber, err)
					}

					return nil
				})
			}(chunkNum)
		}
	}

	if len(chunk) > 0 {
		chunkNum++
		func(orderNumber int) {
			chunkToEncrypt := string(chunk)
			g.Go(func() error {
				encryptedChunk, err := s.Crypt.Encrypt(chunkToEncrypt)
				if err != nil {
					return fmt.Errorf("failed to encrypt chunk %d: %w", orderNumber, err)
				}

				err = s.Storage.AddBinaryChunk(ctx, binID, encryptedChunk, orderNumber)
				if err != nil {
					return fmt.Errorf("failed to add binary chunk %d: %w", orderNumber, err)
				}

				return nil
			})
		}(chunkNum)
	}

	if err := g.Wait(); err != nil {
		return 0, fmt.Errorf("errgroup error: %w", err)
	}

	return binID, nil
}

func (s *Service) AddDataBinaryTitle(ctx context.Context, ownerLogin string, binary *entity.DataBinary) (int32, error) {
	u, err := s.Storage.UserByLogin(ctx, ownerLogin)
	if err != nil {
		return 0, fmt.Errorf("failed to get user by login: %w", err)
	}

	bin, err := s.Storage.BinaryByID(ctx, u.ID, binary.ID)
	if err != nil {
		return 0, fmt.Errorf("failed to get binar by ID: %w", err)
	}

	if err := s.Storage.UpdateBinary(ctx, binary); err != nil {
		return 0, fmt.Errorf("failed to update binary: %w", err)
	}

	return bin.ID, err
}
