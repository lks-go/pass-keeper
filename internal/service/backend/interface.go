package backend

import (
	"context"

	"github.com/lks-go/pass-keeper/internal/service/entity"
)

type Storage interface {
	UserLogin
	RegisterUser(ctx context.Context, login string, passwordHash string) (string, error)

	AddLoginPass(ctx context.Context, owner string, data *entity.DataLoginPass) (int32, error)
	LoginPassList(ctx context.Context, owner string) ([]entity.DataLoginPass, error)
	LoginPassByID(ctx context.Context, owner string, ID int32) (*entity.DataLoginPass, error)

	AddText(ctx context.Context, owner string, data *DataText) (int32, error)
	TextList(ctx context.Context, owner string) ([]DataText, error)
	TextByID(ctx context.Context, owner string, ID int32) (*DataText, error)

	AddCard(ctx context.Context, owner string, data *DataCard) (int32, error)
	CardList(ctx context.Context, owner string) ([]DataCard, error)
	CardByID(ctx context.Context, owner string, ID int32) (*DataCard, error)
}

type UserLogin interface {
	UserByLogin(ctx context.Context, login string) (*entity.User, error)
}

type PasswordHash interface {
	Hash(pass string) string
}

type Crypt interface {
	Encrypt(stringToEncrypt string) (string, error)
	Decrypt(encryptedString string) (string, error)
}
