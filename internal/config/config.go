package config

import (
	"log"

	"github.com/caarlos0/env/v8"
)

type AppConfig struct {
	PgxConnectionString string `env:"POSTGRES_CONNECTION_STRING"`
	PasswordPepper      string `env:"PASSWORD_PEPPER"`
	AccessTokenRoute    string `env:"ACCESS_TOKEN_ROUTE"`
	Domain              string `env:"DOMAIN"`
	PrivateKey          string `env:"PRIVATE_KEY"`
}

func GetConfig() *AppConfig {
	var config AppConfig

	if err := env.Parse(&config); err != nil {
		//nolint:gocritic
		log.Fatal("could not parse config: ", err)
	}

	return &config
}
