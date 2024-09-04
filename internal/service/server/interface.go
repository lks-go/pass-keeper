package server

import "context"

type Storage interface {
	UserLogin
	RegisterUser(ctx context.Context, login string, passwordHash string) (string, error)

	AddLoginPass(ctx context.Context, owner string, data *DataLoginPass) (int32, error)
	LoginPassList(ctx context.Context, owner string) ([]DataLoginPass, error)
	LoginPassByID(ctx context.Context, owner string, ID int32) (*DataLoginPass, error)

	AddText(ctx context.Context, owner string, data *DataText) (int32, error)
	TextList(ctx context.Context, owner string) ([]DataText, error)
	TextByID(ctx context.Context, owner string, ID int32) (*DataText, error)

	AddCard(ctx context.Context, owner string, data *DataCard) (int32, error)
}

type UserLogin interface {
	UserByLogin(ctx context.Context, login string) (*User, error)
}

type PasswordHash interface {
	Hash(pass string) string
}

type Crypt interface {
	Encrypt(stringToEncrypt string) (string, error)
	Decrypt(encryptedString string) (string, error)
}
