package server

import "context"

type Storage interface {
	RegisterUser(ctx context.Context, login string, passwordHash string) (string, error)
	UserByLogin(ctx context.Context, login string) (*User, error)
	AddData(ctx context.Context, owner string, data Data) error
}

type PasswordHash interface {
	Hash(pass string) string
}

type Crypt interface {
	Encrypt(stringToEncrypt string) (string, error)
	Decrypt(encryptedString string) (string, error)
}
