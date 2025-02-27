package config

import (
	env "github.com/caarlos0/env/v6"
	"github.com/charmingruby/impr/service/identity/pkg/logger"
	"github.com/joho/godotenv"
)

func New() (Config, error) {
	if err := godotenv.Load(); err != nil {
		logger.Log.Info(".env file found")
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
		Cognito: cognitoConfig{
			AppClientID: environment.CognitoAppClientID,
			UserPoolID:  environment.CognitoUserPoolID,
		},
	}

	return cfg, nil
}

type environment struct {
	ServerPort         string `env:"SERVER_PORT,required"`
	ServerHost         string `env:"SERVER_HOST,required"`
	CognitoAppClientID string `env:"COGNITO_APP_CLIENT_ID,required"`
	CognitoUserPoolID  string `env:"COGNITO_USER_POLL_ID,required"`
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
	UserPoolID  string
}
