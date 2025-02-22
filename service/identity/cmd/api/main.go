package main

import (
	"fmt"

	"github.com/charmingruby/impr/lib/pkg/awsc"
	"github.com/charmingruby/impr/lib/pkg/rest"
	"github.com/charmingruby/impr/service/identity/config"
	"github.com/charmingruby/impr/service/identity/internal/account/core/gateway"
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

	cognitoClient, err := awsc.NewCognitoClient(cfg.Cognito.AppClientID)
	if err != nil {
		panic(err)
	}

	op, err := client.NewCognitoIdentityProvider(cognitoClient).SignIn(gateway.SignInInput{
		Email:    "gustavodiasa2121@gmail.com",
		Password: "P@ssword123",
	})
	if err != nil {
		panic(err)
	}
	println(op.AccessToken)

	user, err := client.NewCognitoIdentityProvider(cognitoClient).RetrieveUser(op.AccessToken)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", user)

	logger.Log.Info("Connected to Cognito Client.")

	router := echo.New()
	restServer := rest.New(router, cfg.Server.Host, cfg.Server.Port)

	logger.Log.Info(fmt.Sprintf("Rest server is running at: %s:%s ...", cfg.Server.Host, cfg.Server.Port))

	if err := restServer.Start(); err != nil {
		panic(err)
	}
}
