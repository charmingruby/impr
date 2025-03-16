package main

import (
	"fmt"
	"net"
	"os"

	"github.com/charmingruby/impr/lib/pkg/http/client/service/awsc"
	"github.com/charmingruby/impr/lib/pkg/http/server/rest"
	"github.com/charmingruby/impr/service/identity/config"
	"github.com/charmingruby/impr/service/identity/internal/account"
	grpcSvc "github.com/charmingruby/impr/service/identity/internal/account/transport/grpc/server"
	"github.com/charmingruby/impr/service/identity/internal/account/transport/shared/client"
	"github.com/charmingruby/impr/service/identity/pkg/logger"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	logger.New()

	cfg, err := config.New()
	if err != nil {
		logger.Log.Error(err.Error())
		os.Exit(1)
	}

	cognitoCl, err := awsc.NewCognitoClient(cfg.Cognito.AppClientID, cfg.Cognito.UserPoolID)
	if err != nil {
		logger.Log.Error(err.Error())
		os.Exit(1)
	}

	router := echo.New()

	identityProviderClient := client.NewCognitoIdentityProvider(cognitoCl)

	svc := account.NewService(identityProviderClient)

	account.NewRestHandler(router, svc).Register()

	restServer := rest.New(router, cfg.Server.RestHost, cfg.Server.RestPort)

	go func() {
		logger.Log.Info(fmt.Sprintf("Rest server is running at: %s:%s ...", cfg.Server.RestHost, cfg.Server.RestPort))

		if err := restServer.Start(); err != nil {
			logger.Log.Error(err.Error())
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

		grpcSvc.New(server, identityProviderClient).Register()

		reflection.Register(server)

		if err := server.Serve(lis); err != nil {
			logger.Log.Error(err.Error())
			os.Exit(1)
		}
	}()

	select {}
}
