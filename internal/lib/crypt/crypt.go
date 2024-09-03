package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
)

var ErrSecretKeyLen = errors.New("secret key len must be 32 bytes")

func New(secretKey string) (*Crypt, error) {
	if len(secretKey) != 32 {
		return nil, ErrSecretKeyLen
	}

	return &Crypt{secretKey: []byte(secretKey)}, nil
}

type Crypt struct {
	secretKey []byte
}

func (c *Crypt) Encrypt(stringToEncrypt string) (string, error) {
	block, err := aes.NewCipher(c.secretKey)
	if err != nil {
		return "", fmt.Errorf("filed to create new cipher block: %w", err)
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("filed to create new GCM: %w", err)
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", fmt.Errorf("filed to make nonce: %w", err)
	}

	ciphertext := aesGCM.Seal(nonce, nonce, []byte(stringToEncrypt), nil)

	return fmt.Sprintf("%x", ciphertext), nil
}

func (c *Crypt) Decrypt(encryptedString string) (string, error) {
	enc, err := hex.DecodeString(encryptedString)
	if err != nil {
		return "", fmt.Errorf("failed to decode encrypted string: %w", err)
	}

	block, err := aes.NewCipher(c.secretKey)
	if err != nil {
		return "", fmt.Errorf("filed to create new cipher block: %w", err)
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("filed to create new GCM: %w", err)
	}

	nonceSize := aesGCM.NonceSize()
	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", fmt.Errorf("filed to open cipher text: %w", err)
	}

	return fmt.Sprintf("%s", plaintext), nil
}
