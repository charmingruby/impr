package main

import (
	"github.com/charmingruby/impr/service/identity/config"
	"github.com/charmingruby/impr/service/identity/internal/shared/transport/rest"
	"github.com/labstack/echo/v4"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	router := echo.New()

	restServer := rest.New(router, cfg.Server.Host, cfg.Server.Port)

	println("server is running...")

	if err := restServer.Start(); err != nil {
		panic(err)
	}
}
