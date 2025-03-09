package main

import (
	"fmt"
	"os"

	"github.com/charmingruby/impr/lib/pkg/http/server/rest"
	"github.com/charmingruby/impr/service/audit/pkg/logger"
	"github.com/charmingruby/impr/service/poll/config"
	"github.com/charmingruby/impr/service/poll/internal/poll"
	"github.com/charmingruby/impr/service/poll/pkg/postgres"
	"github.com/charmingruby/impr/service/poll/test/memory"
	"github.com/labstack/echo/v4"
)

func main() {
	logger.New()

	cfg, err := config.New()
	if err != nil {
		logger.Log.Error(err.Error())
		os.Exit(1)
	}

	logger.Log.Info("Connecting to Postgres...")

	_, err = postgres.New(postgres.ConnectionInput{
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

	pollRepo := memory.NewPollRepository()
	optionRepo := memory.NewPollOptionRepository()
	voteRepo := memory.NewVoteRepository()

	svc := poll.NewService(pollRepo, optionRepo, voteRepo)

	router := echo.New()

	poll.NewRestHandler(router, svc).Register()

	restServer := rest.New(router, cfg.Server.Host, cfg.Server.Port)

	logger.Log.Info(fmt.Sprintf("Rest server is running at: %s:%s ...", cfg.Server.Host, cfg.Server.Port))

	if err := restServer.Start(); err != nil {
		os.Exit(1)
	}
}
