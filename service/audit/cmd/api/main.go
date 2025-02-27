package main

import (
	"github.com/charmingruby/impr/service/audit/config"
	"github.com/charmingruby/impr/service/audit/pkg/logger"
	"github.com/charmingruby/impr/service/audit/pkg/mongodb"
)

func main() {
	logger.New()

	cfg, err := config.New()
	if err != nil {
		logger.Log.Error(err.Error())
		panic(err)
	}

	logger.Log.Info("Connecting to Mongo...")

	_, err = mongodb.New(cfg.Mongo.MongoURL, cfg.Mongo.MongoDatabase)
	if err != nil {
		logger.Log.Error(err.Error())
		panic(err)
	}

	logger.Log.Info("Connected to Mongo successfully!")
}
