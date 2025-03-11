package endpoint

import (
	"github.com/charmingruby/impr/service/poll/internal/poll/core/service"
	"github.com/labstack/echo/v4"
)

type Endpoint struct {
	r *echo.Echo

	service *service.Service
}

func New(
	r *echo.Echo,
	svc *service.Service,
) *Endpoint {
	return &Endpoint{
		r:       r,
		service: svc,
	}
}

func (e *Endpoint) Register() {
	api := e.r.Group("/api")

	api.POST("/poll", e.makeCreatePollEndpoint())
	api.GET("/poll/:poll_id", e.makeGetPollDetailsEndpoint())
	api.PATCH("/poll/:poll_id/close", e.makeClosePollEndpoint())
	api.POST("/poll/:poll_id/vote", e.makeVoteOnPollEndpoint())
}
