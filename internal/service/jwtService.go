package service

import (
	"fmt"
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
		Subject:   fmt.Sprint(user.Id),
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

func (service *JwtService) ValidateToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(service.config.PrivateKey), nil
	})

	return nil

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["foo"], claims["nbf"])
	} else {
		fmt.Println(err)
	}

	if err != nil {
		return fmt.Errorf("Invalid token - %w", err)
	}

	return nil
}
