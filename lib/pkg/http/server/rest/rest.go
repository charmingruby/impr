package rest

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type Server struct {
	HTTPServer *http.Server
}

func New(router *echo.Echo, host string, port string) *Server {
	addr := fmt.Sprintf("%s:%s", host, port)

	server := http.Server{
		Addr:         addr,
		Handler:      router,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}

	return &Server{
		HTTPServer: &server,
	}
}

func (s *Server) Start() error {
	return s.HTTPServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.HTTPServer.Shutdown(ctx)
}
