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

	AddText(ctx context.Context, owner string, data *entity.DataText) (int32, error)
	TextList(ctx context.Context, owner string) ([]entity.DataText, error)
	TextByID(ctx context.Context, owner string, ID int32) (*entity.DataText, error)

	AddCard(ctx context.Context, owner string, data *entity.DataCard) (int32, error)
	CardList(ctx context.Context, owner string) ([]entity.DataCard, error)
	CardByID(ctx context.Context, owner string, ID int32) (*entity.DataCard, error)

	AddBinary(ctx context.Context, owner string, binary *entity.DataBinary) (int32, error)
	UpdateBinary(ctx context.Context, binary *entity.DataBinary) error
	BinaryList(ctx context.Context, owner string) ([]entity.DataBinary, error)
	BinaryByID(ctx context.Context, owner string, ID int32) (*entity.DataBinary, error)
	AddBinaryChunk(ctx context.Context, binaryID int32, encryptedData string, orderNumber int) error
	BinaryChunk(ctx context.Context, binaryID int32, orderNumber int) (encryptedData string, err error)
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
