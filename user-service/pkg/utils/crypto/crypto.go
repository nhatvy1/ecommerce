package crypto

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

func GenerateSalt(length int) (string, error) {
	salt := make([]byte, length)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	return hex.EncodeToString(salt), nil
}

func HashPassword(password string, salt string) string {
	saltedPassword := password + salt
	hashPassword := sha256.Sum256(([]byte(saltedPassword)))

	return hex.EncodeToString(hashPassword[:])
}

func MatchingPassword(storeHash string, password string, salt string) bool {
	HashPassword := HashPassword(password, salt)
	return storeHash == HashPassword
}
