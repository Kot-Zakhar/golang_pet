package service

import (
	"crypto/rand"
	"crypto/sha512"

	"github.com/kot-zakhar/golang_pet/internal/config"
)

type PasswordService struct {
	config *config.AppConfig
}

// Q: should Factory return interface or concrete class?
// potential A: Class, because interfaces are partial
// Q: how to state inside of current file, that the class should implement a list of interfaces?
// Otherwise the mistakes are going to be in usage section, which is missleading
// current PasswordService implements two interfaces: IPasswordHasherService and IPasswordCheckerService
func NewPasswordService(config *config.AppConfig) PasswordService {
	return PasswordService{config}
}

func (service *PasswordService) HashPassword(password string) (hash, salt []byte) {
	// Yeah, normal guys use bcrypt
	// but this is just for practice.
	salt = service.generateSalt()
	pepper := service.getPepper()
	hash = service.hashPasswordWithSaltAndPepper(password, salt, pepper)
	return
}

func (service *PasswordService) DoPasswordsMatch(currPassword string, hashedPassword, salt []byte) bool {
	pepper := service.getPepper()
	currPasswordHash := service.hashPasswordWithSaltAndPepper(currPassword, salt, pepper)

	return string(hashedPassword) == string(currPasswordHash)
}

const saltSize = 16

func (service *PasswordService) generateSalt() (salt []byte) {
	salt = make([]byte, saltSize)

	_, err := rand.Read(salt)

	if err != nil {
		panic(err)
	}

	return
}

func (service *PasswordService) getPepper() []byte {
	return []byte(service.config.PasswordPepper)
}

func (service *PasswordService) hashPasswordWithSaltAndPepper(password string, salt, pepper []byte) []byte {
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
