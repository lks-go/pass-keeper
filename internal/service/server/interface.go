package server

import "context"

type Storage interface {
	UserLogin
	RegisterUser(ctx context.Context, login string, passwordHash string) (string, error)
	AddLoginPass(ctx context.Context, owner string, data LoginPassData) error
	LoginPassList(ctx context.Context, owner string) ([]LoginPassData, error)
	LoginPassByID(ctx context.Context, owner string, ID int32) (*LoginPassData, error)
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
