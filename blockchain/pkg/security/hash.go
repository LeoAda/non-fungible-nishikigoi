package security

import (
	"crypto/sha256"
	"encoding/hex"
)

// HashString : hash a string in sha256
func HashString(s string) (string, error) {
	h := sha256.New()
	_, err := h.Write([]byte(s))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
