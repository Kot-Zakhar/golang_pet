package service

import (
	"crypto/rand"
	"crypto/sha512"

	"github.com/kot-zakhar/golang_pet/internal/config"
)

type CryptoService struct {
	config *config.AppConfig
}

func NewCryptoService(config *config.AppConfig) CryptoService {
	return CryptoService{config}
}

func (service *CryptoService) HashPassword(password string) (hash, salt []byte) {
	// Yeah, normal guys use bcrypt
	// but this is just for practice.
	salt = service.generateSalt()
	pepper := service.getPepper()
	hash = service.hashPassword(password, salt, pepper)
	return
}

func (service *CryptoService) DoPasswordsMatch(currPassword string, hashedPassword, salt []byte) bool {
	pepper := service.getPepper()
	currPasswordHash := service.hashPassword(currPassword, salt, pepper)
	return string(hashedPassword) == string(currPasswordHash)
}

const saltSize = 16

func (service *CryptoService) generateSalt() (salt []byte) {
	salt = make([]byte, saltSize)

	_, err := rand.Read(salt)

	if err != nil {
		panic(err)
	}

	return
}

func (service *CryptoService) getPepper() []byte {
	return []byte(service.config.PasswordPepper)
}

func (service *CryptoService) hashPassword(password string, salt, pepper []byte) []byte {
	pepperBytes := []byte(pepper)
	saltBytes := []byte(salt)
	passwordBytes := []byte(password)
	sha512Hasher := sha512.New()

	passwordBytes = append(passwordBytes, saltBytes...)

	sha512Hasher.Write(passwordBytes)

	pepperedPasswordBytes := sha512Hasher.Sum(pepperBytes)

	sha512Hasher.Reset()
	sha512Hasher.Write(pepperedPasswordBytes)

	return sha512Hasher.Sum(nil)
}
