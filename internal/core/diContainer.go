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
}

var DI DiContainer

func InitializeDI(config *config.AppConfig) {
	dbConnection, err := ConnectPgx(config.PgxConnectionString)
	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := repository.NewUserRepository(dbConnection)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	DI = DiContainer{
		UserHandler: userHandler,
	}
}
