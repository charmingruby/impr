package config

import (
	"log"

	env "github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

func New() (Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Println(".env file found")
	}

	environment := environment{}
	if err := env.Parse(&environment); err != nil {
		return Config{}, err
	}

	cfg := Config{
		Server: serverConfig{
			Port: environment.ServerPort,
			Host: environment.ServerHost,
		},
	}

	return cfg, nil
}

type environment struct {
	ServerPort         string `env:"SERVER_PORT,required"`
	ServerHost         string `env:"SERVER_HOST,required"`
	CognitoAppClientID string `env:"COGNITO_APP_CLIENT_ID,required"`
}

type Config struct {
	Server  serverConfig
	Cognito cognitoConfig
}

type serverConfig struct {
	Host string
	Port string
}

type cognitoConfig struct {
	AppClientID string
}
