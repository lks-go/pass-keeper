package crypt_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/lks-go/pass-keeper/internal/lib/crypt"
)

func TestEncryptDecrypt(t *testing.T) {
	testKey := "KrfD58EwwCo9e1YTqaGfg66n3skI90qX"
	stringToEncrypt := "some text"

	c, err := crypt.New(testKey)
	require.NoError(t, err)

	encryptedString, err := c.Encrypt(stringToEncrypt)
	require.NoError(t, err)

	decryptedString, err := c.Decrypt(encryptedString)
	require.NoError(t, err)

	require.Equal(t, stringToEncrypt, decryptedString)
}

func TestSecretKeyLenError(t *testing.T) {
	testKey := "KrfD58EwwCo9e1YTqaGfg66n3skI90q"

	_, err := crypt.New(testKey)
	require.ErrorIs(t, err, crypt.ErrSecretKeyLen)
}
