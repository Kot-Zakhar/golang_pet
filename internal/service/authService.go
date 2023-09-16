package service

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/kot-zakhar/golang_pet/internal/config"
	"github.com/kot-zakhar/golang_pet/internal/model"
)

const RefreshTokenAge = 60 * 24 * time.Hour

type ILoginUserRepository interface {
	GetByLogin(context.Context, string) (model.User, error)
}

type IAuthRepository interface {
	InsertSession(context.Context, model.UserSession) (model.UserSession, error)
}

type IPasswordCheckerService interface {
	DoPasswordsMatch(currPassword string, hashedPassword, salt []byte) bool
}

type IJwtService interface {
	CreateToken(user model.User, session model.UserSession) (string, error)
	ValidateToken(tokenString string) error
}

type AuthService struct {
	config                 *config.AppConfig
	authRepository         IAuthRepository
	userRepository         ILoginUserRepository
	passowrdCheckerService IPasswordCheckerService
	jwtService             IJwtService
}

func NewAuthService(
	config *config.AppConfig,
	authRepository IAuthRepository,
	userRepository ILoginUserRepository,
	passwordCheckerService IPasswordCheckerService,
	jwtService IJwtService,
) AuthService {
	return AuthService{
		config,
		authRepository,
		userRepository,
		passwordCheckerService,
		jwtService,
	}
}

func (service *AuthService) SignIn(context context.Context,
	login string, password string, fingerprint string, userAgent string,
) (accessToken string, session model.UserSession, err error) {

	user, err := service.userRepository.GetByLogin(context, login)
	if err != nil {
		err = fmt.Errorf("AuthService:SignIn - %w", err)
		return
	}

	if !service.passowrdCheckerService.DoPasswordsMatch(password, user.PasswordHash, user.Salt) {
		err = fmt.Errorf("Incorrect password")
		return
	}

	// TODO: remove existing sessions if more than 5 open sessions exist

	session = model.UserSession{
		UserId:       user.Id,
		UserAgent:    userAgent,
		RefreshToken: uuid.New().String(),
		Fingerprint:  fingerprint,
		CreatedAt:    time.Now(),
		ExpiresAt:    time.Now().Add(RefreshTokenAge),
	}

	session, err = service.authRepository.InsertSession(context, session)
	if err != nil {
		err = fmt.Errorf("AuthService:SignIn.InsertSession - %w", err)
		return
	}

	accessToken, err = service.jwtService.CreateToken(user, session)
	if err != nil {
		err = fmt.Errorf("AuthService:SignIn.CreateToken - %w", err)
		return
	}

	return
}

func (service *AuthService) SignOut(context context.Context, login string) error {
	return nil
}

func (service *AuthService) RefreshTokens(context context.Context, oldRefreshToken string) (accessToken, newRefreshToken string, err error) {
	return "", "", nil
}
