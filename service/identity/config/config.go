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
			RestPort: environment.RestServerPort,
			RestHost: environment.RestServerHost,
			GRPCPort: environment.GRPCServerPort,
			GRPCHost: environment.GRPCServerHost,
		},
		Cognito: cognitoConfig{
			AppClientID: environment.CognitoAppClientID,
			UserPoolID:  environment.CognitoUserPoolID,
		},
	}

	return cfg, nil
}

type environment struct {
	RestServerPort     string `env:"REST_SERVER_PORT,required"`
	RestServerHost     string `env:"REST_SERVER_HOST,required"`
	GRPCServerPort     string `env:"GRPC_SERVER_PORT,required"`
	GRPCServerHost     string `env:"GRPC_SERVER_HOST,required"`
	CognitoAppClientID string `env:"COGNITO_APP_CLIENT_ID,required"`
	CognitoUserPoolID  string `env:"COGNITO_USER_POLL_ID,required"`
}

type Config struct {
	Server  serverConfig
	Cognito cognitoConfig
}

type serverConfig struct {
	RestHost string
	RestPort string
	GRPCHost string
	GRPCPort string
}

type cognitoConfig struct {
	AppClientID string
	UserPoolID  string
}
