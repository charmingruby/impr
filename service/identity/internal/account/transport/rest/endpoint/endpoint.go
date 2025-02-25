package endpoint

import (
	"github.com/charmingruby/impr/service/identity/internal/account/core/service"
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
	g := e.r.Group("/api/auth")

	g.POST("/signup", e.makeSignUpEndpoint())
	g.POST("/confirm-account", e.makeConfirmAccountEndpoint())
	g.POST("/signin", e.makeSignInEndpoint())
}
