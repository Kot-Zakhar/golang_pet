package service

import (
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"

	"github.com/kot-zakhar/golang_pet/internal/config"
	"github.com/kot-zakhar/golang_pet/internal/model"
)

const tokenAge = 30 * time.Minute

type JwtService struct {
	config *config.AppConfig
}

func NewJwtService(config *config.AppConfig) JwtService {
	return JwtService{config}
}

func (service *JwtService) CreateToken(user model.User, session model.UserSession) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Subject:   strconv.Itoa(user.Id),
		Issuer:    service.config.Domain,
		Audience:  session.Fingerprint,
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(tokenAge).Unix(),
	})

	signedToken, err := token.SignedString([]byte(service.config.PrivateKey))

	if err != nil {
		return "", fmt.Errorf("Error while signing JWT token - %w", err)
	}

	return signedToken, nil
}

func (service *JwtService) ValidateAndGetUserId(tokenString string) (userId string, err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(service.config.PrivateKey), nil
	})

	if err != nil {
		return "", fmt.Errorf("Error parsing token - %w", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId = claims["sub"].(string)
	} else {
		fmt.Println("Invalid signature")
		err = fmt.Errorf("Signature is not valid")
	}

	return
}
