package main

import (
	"fmt"

	"github.com/charmingruby/impr/lib/pkg/client/service/awsc"
	"github.com/charmingruby/impr/lib/pkg/server/rest"
	"github.com/charmingruby/impr/service/identity/config"
	"github.com/charmingruby/impr/service/identity/internal/account"
	"github.com/charmingruby/impr/service/identity/internal/account/transport/rest/client"
	"github.com/charmingruby/impr/service/identity/pkg/logger"
	"github.com/labstack/echo/v4"
)

func main() {
	logger.New()

	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	cognitoCl, err := awsc.NewCognitoClient(cfg.Cognito.AppClientID, cfg.Cognito.UserPoolID)
	if err != nil {
		panic(err)
	}

	router := echo.New()

	identityProviderClient := client.NewCognitoIdentityProvider(cognitoCl)

	accountSvc := account.NewService(identityProviderClient)

	account.NewRestHandler(router, accountSvc).Register()

	restServer := rest.New(router, cfg.Server.Host, cfg.Server.Port)

	logger.Log.Info(fmt.Sprintf("Rest server is running at: %s:%s ...", cfg.Server.Host, cfg.Server.Port))

	if err := restServer.Start(); err != nil {
		panic(err)
	}
}
