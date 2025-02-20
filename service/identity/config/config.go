package config

import (
	"log"

	env "github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	Server serverConfig
}

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
