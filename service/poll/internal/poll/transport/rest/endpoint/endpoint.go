package endpoint

import (
	"github.com/charmingruby/impr/service/poll/internal/poll/core/service"
	"github.com/charmingruby/impr/service/poll/internal/shared/transport/rest/middleware"
	"github.com/labstack/echo/v4"
)

type Endpoint struct {
	r          *echo.Echo
	service    *service.Service
	middleware *middleware.Auth
}

func New(
	r *echo.Echo,
	svc *service.Service,
	mw *middleware.Auth,
) *Endpoint {
	return &Endpoint{
		r:          r,
		service:    svc,
		middleware: mw,
	}
}

func (e *Endpoint) Register() {
	api := e.r.Group("/api")

	api.POST("/poll", e.middleware.Intercept(e.makeCreatePollEndpoint()))
	api.GET("/poll/:poll_id", e.makeGetPollDetailsEndpoint())
	api.PATCH("/poll/:poll_id/close", e.middleware.Intercept(e.makeClosePollEndpoint()))
	api.POST("/poll/:poll_id/vote", e.middleware.Intercept(e.makeVoteOnPollEndpoint()))
}
