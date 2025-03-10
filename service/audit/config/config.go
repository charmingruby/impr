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
		Mongo: mongoConfig{
			MongoURL:      environment.MongoURL,
			MongoDatabase: environment.MongoDatabase,
		},
		Kafka: kafkaConfig{
			BrokerURL:        environment.KafkaBrokerURL,
			GroupID:          environment.KafkaGroupID,
			CreateAuditTopic: environment.KafkaCreateAuditTopic,
		},
	}

	return cfg, nil
}

type environment struct {
	ServerHost            string `env:"SERVER_HOST,required"`
	ServerPort            string `env:"SERVER_PORT,required"`
	MongoURL              string `env:"MONGO_URL,required"`
	MongoDatabase         string `env:"MONGO_DB,required"`
	KafkaBrokerURL        string `env:"KAFKA_BROKER_URL,required"`
	KafkaGroupID          string `env:"KAFKA_GROUP_ID,required"`
	KafkaCreateAuditTopic string `env:"KAFKA_CREATE_AUDIT_TOPIC,required"`
}

type Config struct {
	Server serverConfig
	Mongo  mongoConfig
	Kafka  kafkaConfig
}

type serverConfig struct {
	Host string
	Port string
}

type mongoConfig struct {
	MongoURL      string
	MongoDatabase string
}

type kafkaConfig struct {
	BrokerURL string
	GroupID   string

	CreateAuditTopic string
}
