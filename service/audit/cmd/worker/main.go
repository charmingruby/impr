package main

import (
	"context"

	"github.com/charmingruby/impr/lib/pkg/messaging/kafka"
	"github.com/charmingruby/impr/service/audit/config"
	"github.com/charmingruby/impr/service/audit/internal/audit/core/event"
	"github.com/charmingruby/impr/service/audit/internal/audit/core/service"
	"github.com/charmingruby/impr/service/audit/internal/audit/database/mongodb"
	"github.com/charmingruby/impr/service/audit/pkg/logger"
	mongoConn "github.com/charmingruby/impr/service/audit/pkg/mongodb"
)

func main() {
	logger.New()

	cfg, err := config.New()
	if err != nil {
		logger.Log.Error(err.Error())
		panic(err)
	}

	logger.Log.Info("Connecting to Mongo...")

	db, err := mongoConn.New(cfg.Mongo.MongoURL, cfg.Mongo.MongoDatabase)
	if err != nil {
		logger.Log.Error(err.Error())
		panic(err)
	}

	logger.Log.Info("Connected to Mongo successfully!")

	repo := mongodb.NewAuditRepository(db)

	svc := service.New(repo)

	logger.Log.Info("Connection to Kafka...")

	createAuditSubscriber, err := kafka.NewSubscriber(cfg.Kafka.BrokerURL, cfg.Kafka.CreateAuditTopic, cfg.Kafka.GroupID)
	if err != nil {
		logger.Log.Error(err.Error())
		panic(err)
	}
	defer createAuditSubscriber.Close()

	ctx := context.Background()
	go event.HandleCreateAudit(ctx, createAuditSubscriber, svc)

	logger.Log.Info("Listening to Kafka messages...")

	select {}
}
