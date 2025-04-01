package crypto

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"user-service/global"
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

func GetHash(key string) string {
	salt := global.Config.Salt
	hash := sha256.New()
	hash.Write([]byte(key + salt))
	hashBytes := hash.Sum(nil)
	return hex.EncodeToString(hashBytes)
}

func GetHashNoSalt(key string) string {
	hash := sha256.New()
	hash.Write([]byte(key))
	hashBytes := hash.Sum(nil)
	return hex.EncodeToString(hashBytes)
}
