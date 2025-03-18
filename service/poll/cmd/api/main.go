package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/charmingruby/impr/lib/pkg/http/server/rest"
	"github.com/charmingruby/impr/lib/pkg/messaging/kafka"
	"github.com/charmingruby/impr/service/poll/config"
	"github.com/charmingruby/impr/service/poll/internal/poll"
	"github.com/charmingruby/impr/service/poll/internal/poll/core/service"
	grpcSvc "github.com/charmingruby/impr/service/poll/internal/poll/transport/grpc/server"
	"github.com/charmingruby/impr/service/poll/internal/shared/transport/grpc/client"
	"github.com/charmingruby/impr/service/poll/internal/shared/transport/rest/middleware"
	"github.com/charmingruby/impr/service/poll/pkg/logger"
	"github.com/charmingruby/impr/service/poll/pkg/postgres"
	"github.com/labstack/echo/v4"
	"github.com/robfig/cron/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
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

	summaryRepo, err := poll.NewPollSumaryRepository(db)
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

	svc := poll.NewService(pollRepo, optionRepo, voteRepo, summaryRepo, publisher)

	closeExpiredPollsCronJob(30, svc)

	router := echo.New()

	authMiddleware := middleware.NewAuth(gRPCClient.Service)

	poll.NewRestHandler(router, svc, authMiddleware).Register()

	restServer := rest.New(router, cfg.Server.RestHost, cfg.Server.RestPort)

	go func() {
		logger.Log.Info(fmt.Sprintf("Rest server is running at: %s:%s ...", cfg.Server.RestHost, cfg.Server.RestPort))

		if err := restServer.Start(); err != nil {
			os.Exit(1)
		}
	}()

	go func() {
		logger.Log.Info(fmt.Sprintf("gRPC server is running at: %s:%s ...", cfg.Server.GRPCHost, cfg.Server.GRPCPort))

		grpcAddr := fmt.Sprintf("%s:%s", cfg.Server.GRPCHost, cfg.Server.GRPCPort)

		lis, err := net.Listen("tcp", grpcAddr)
		if err != nil {
			logger.Log.Error(err.Error())
			os.Exit(1)
		}

		server := grpc.NewServer()

		grpcSvc.New(server, svc).Register()

		reflection.Register(server)

		if err := server.Serve(lis); err != nil {
			logger.Log.Error(err.Error())
			os.Exit(1)
		}
	}()

	select {}
}

func closeExpiredPollsCronJob(delayInSeconds int, svc *service.Service) {
	c := cron.New(cron.WithSeconds())

	c.AddFunc(fmt.Sprintf("*/%d * * * * *", delayInSeconds), func() {
		logger.Log.Info(fmt.Sprintf("CronJob: Closing expired polls at %s", time.Now().String()))

		if err := svc.CloseExpiredPolls(); err != nil {
			logger.Log.Error(err.Error())
		}
	})

	c.Start()
}
