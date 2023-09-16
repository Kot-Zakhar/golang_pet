package core

import (
	"log"

	"github.com/kot-zakhar/golang_pet/internal/config"
	"github.com/kot-zakhar/golang_pet/internal/handler"
	"github.com/kot-zakhar/golang_pet/internal/repository"
	"github.com/kot-zakhar/golang_pet/internal/service"
)

// acts as a buch of singletons for the moment

type DiContainer struct {
	// handlers
	UserHandler handler.UserHandler
	AuthHandler handler.AuthHandler
}

var DI DiContainer

func InitializeDI(config *config.AppConfig) {
	dbConnection, err := ConnectPgx(config.PgxConnectionString)
	if err != nil {
		log.Fatal(err.Error())
	}

	passwordService := service.NewPasswordService(config)
	jwtService := service.NewJwtService(config)

	userRepository := repository.NewUserRepository(dbConnection)
	userService := service.NewUserService(&userRepository, &passwordService)
	userHandler := handler.NewUserHandler(&userService)

	authRepository := repository.NewAuthRepository(dbConnection)
	authService := service.NewAuthService(config, &authRepository, &userRepository, &passwordService, &jwtService)
	authHandler := handler.NewAuthHandler(config, &authService)

	DI = DiContainer{
		userHandler,
		authHandler,
	}
}
