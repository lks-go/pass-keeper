package password

import (
	"crypto/sha256"
	"encoding/hex"
)

func New(salt string) *Password {
	return &Password{salt: salt}
}

type Password struct {
	salt string
}

func (h *Password) Hash(pass string) string {
	sha := sha256.New()
	sha.Write([]byte(pass + h.salt))
	hash := sha.Sum(nil)
	return hex.EncodeToString(hash)
}
