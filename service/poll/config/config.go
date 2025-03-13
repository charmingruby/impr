package config

import (
	"log/slog"

	env "github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type environment struct {
	ServerPort             string `env:"SERVER_PORT,required"`
	ServerHost             string `env:"SERVER_HOST,required"`
	IdentityGRPCServerPort string `env:"IDENTITY_GRPC_SERVER_PORT,required"`
	IdentityGRPCServerHost string `env:"IDENTITY_GRPC_SERVER_HOST,required"`
	DatabaseUser           string `env:"DATABASE_USER,required"`
	DatabasePassword       string `env:"DATABASE_PASSWORD,required"`
	DatabaseHost           string `env:"DATABASE_HOST,required"`
	DatabaseName           string `env:"DATABASE_NAME,required"`
	DatabaseSSL            string `env:"DATABASE_SSL,required"`
}

func New() (Config, error) {
	if err := godotenv.Load(); err != nil {
		slog.Warn("CONFIGURATION: .env file not found")
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
		IdentityIntegration: identityIntegrationConfig{
			Port: environment.IdentityGRPCServerPort,
			Host: environment.IdentityGRPCServerHost,
		},
		Postgres: postgresConfig{
			User:         environment.DatabaseUser,
			Password:     environment.DatabasePassword,
			Host:         environment.DatabaseHost,
			DatabaseName: environment.DatabaseName,
			SSL:          environment.DatabaseSSL,
		},
	}

	return cfg, nil
}

type Config struct {
	Server              serverConfig
	IdentityIntegration identityIntegrationConfig
	Postgres            postgresConfig
}

type serverConfig struct {
	Host string
	Port string
}

type identityIntegrationConfig struct {
	Host string
	Port string
}

type postgresConfig struct {
	User         string
	Password     string
	Host         string
	DatabaseName string
	SSL          string
}
