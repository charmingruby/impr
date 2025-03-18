package config

import (
	"log/slog"

	env "github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type environment struct {
	RestServerPort         string `env:"REST_SERVER_PORT,required"`
	RestServerHost         string `env:"REST_SERVER_HOST,required"`
	GRPCServerPort         string `env:"GRPC_SERVER_PORT,required"`
	GRPCServerHost         string `env:"GRPC_SERVER_HOST,required"`
	IdentityGRPCServerPort string `env:"IDENTITY_GRPC_SERVER_PORT,required"`
	IdentityGRPCServerHost string `env:"IDENTITY_GRPC_SERVER_HOST,required"`
	DatabaseUser           string `env:"DATABASE_USER,required"`
	DatabasePassword       string `env:"DATABASE_PASSWORD,required"`
	DatabaseHost           string `env:"DATABASE_HOST,required"`
	DatabaseName           string `env:"DATABASE_NAME,required"`
	DatabaseSSL            string `env:"DATABASE_SSL,required"`
	KafkaBrokerURL         string `env:"KAFKA_BROKER_URL,required"`
	KafkaGroupID           string `env:"KAFKA_GROUP_ID,required"`
	KafkaCreateAuditTopic  string `env:"KAFKA_CREATE_AUDIT_TOPIC,required"`
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
			RestHost: environment.RestServerHost,
			RestPort: environment.RestServerPort,
			GRPCHost: environment.GRPCServerHost,
			GRPCPort: environment.GRPCServerPort,
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
		Kafka: kafkaConfig{
			BrokerURL:        environment.KafkaBrokerURL,
			GroupID:          environment.KafkaGroupID,
			CreateAuditTopic: environment.KafkaCreateAuditTopic,
		},
	}

	return cfg, nil
}

type Config struct {
	Server              serverConfig
	IdentityIntegration identityIntegrationConfig
	Postgres            postgresConfig
	Kafka               kafkaConfig
}

type serverConfig struct {
	RestHost string
	RestPort string
	GRPCHost string
	GRPCPort string
}

type identityIntegrationConfig struct {
	Host string
	Port string
}

type kafkaConfig struct {
	BrokerURL string
	GroupID   string

	CreateAuditTopic string
}

type postgresConfig struct {
	User         string
	Password     string
	Host         string
	DatabaseName string
	SSL          string
}
