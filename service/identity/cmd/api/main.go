package main

import (
	"fmt"

	"github.com/charmingruby/impr/service/identity/config"
	"github.com/charmingruby/impr/service/identity/internal/shared/transport/rest"
	"github.com/charmingruby/impr/service/identity/pkg/aws"
	"github.com/charmingruby/impr/service/identity/pkg/logger"
	"github.com/labstack/echo/v4"
)

func main() {
	logger.New()

	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	_, err = aws.NewCognitoClient(cfg.Cognito.AppClientID)
	if err != nil {
		panic(err)
	}

	logger.Log.Info("Connected to Cognito Client.")

	router := echo.New()
	restServer := rest.New(router, cfg.Server.Host, cfg.Server.Port)

	logger.Log.Info(fmt.Sprintf("Rest server is running at: %s:%s ...", cfg.Server.Host, cfg.Server.Port))

	if err := restServer.Start(); err != nil {
		panic(err)
	}
}
