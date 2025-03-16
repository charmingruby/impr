package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/charmingruby/impr/lib/pkg/http/server/rest"
	"github.com/charmingruby/impr/lib/pkg/messaging/kafka"
	"github.com/charmingruby/impr/service/poll/config"
	"github.com/charmingruby/impr/service/poll/internal/poll"
	"github.com/charmingruby/impr/service/poll/internal/shared/transport/grpc/client"
	"github.com/charmingruby/impr/service/poll/internal/shared/transport/rest/middleware"
	"github.com/charmingruby/impr/service/poll/pkg/logger"
	"github.com/charmingruby/impr/service/poll/pkg/postgres"
	"github.com/labstack/echo/v4"
	"github.com/robfig/cron/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	logger.New()

	cfg, err := config.New()
	if err != nil {
		logger.Log.Error(err.Error())
		os.Exit(1)
	}

	logger.Log.Info(fmt.Sprintf("Connecting to gRPC Identity server at: %s:%s ...", cfg.IdentityIntegration.Host, cfg.IdentityIntegration.Port))

	identityGRPCServerAddr := fmt.Sprintf("%s:%s", cfg.IdentityIntegration.Host, cfg.IdentityIntegration.Port)

	gRPCConn, err := grpc.NewClient(identityGRPCServerAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Unable to connect to server: %v", err)
	}
	defer gRPCConn.Close()

	gRPCClient := client.New(gRPCConn)
	gRPCClient.Register()

	logger.Log.Info("Connected to gRPC Identity server successfully!")

	logger.Log.Info("Connecting to Postgres...")

	db, err := postgres.New(postgres.ConnectionInput{
		User:         cfg.Postgres.User,
		Password:     cfg.Postgres.Password,
		Host:         cfg.Postgres.Host,
		DatabaseName: cfg.Postgres.DatabaseName,
		SSL:          cfg.Postgres.SSL,
	})
	if err != nil {
		logger.Log.Error(err.Error())

		os.Exit(1)
	}

	logger.Log.Info("Connected to Postgres successfully!")

	pollRepo, err := poll.NewPollRepository(db)
	if err != nil {
		logger.Log.Error(err.Error())

		os.Exit(1)
	}

	optionRepo, err := poll.NewPollOptionRepository(db)
	if err != nil {
		logger.Log.Error(err.Error())

		os.Exit(1)
	}

	voteRepo, err := poll.NewVoteRepository(db)
	if err != nil {
		logger.Log.Error(err.Error())

		os.Exit(1)
	}

	logger.Log.Info("Connection to Kafka...")

	publisher, err := kafka.NewPublisher(cfg.Kafka.BrokerURL, cfg.Kafka.CreateAuditTopic)
	if err != nil {
		logger.Log.Error(err.Error())

		os.Exit(1)
	}

	logger.Log.Info("Connected to Kafka successfully!")

	svc := poll.NewService(pollRepo, optionRepo, voteRepo, publisher)

	c := cron.New(cron.WithSeconds())
	c.AddFunc("*/30 * * * * *", func() {
		logger.Log.Info(fmt.Sprintf("CronJob: Closing expired polls at %s", time.Now().String()))

		if err := svc.CloseExpiredPolls(); err != nil {
			logger.Log.Error(err.Error())
		}
	})
	c.Start()

	router := echo.New()

	authMiddleware := middleware.NewAuth(gRPCClient.Service)

	poll.NewRestHandler(router, svc, authMiddleware).Register()

	restServer := rest.New(router, cfg.Server.Host, cfg.Server.Port)

	go func() {
		logger.Log.Info(fmt.Sprintf("Rest server is running at: %s:%s ...", cfg.Server.Host, cfg.Server.Port))

		if err := restServer.Start(); err != nil {
			os.Exit(1)
		}
	}()

	select {}
}
