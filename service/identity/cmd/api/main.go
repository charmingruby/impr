package main

import (
	"github.com/charmingruby/impr/service/gateway/config"
	"github.com/charmingruby/impr/service/gateway/internal/transport/rest"
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
