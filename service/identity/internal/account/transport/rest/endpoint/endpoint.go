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
	e.r.POST("/signup", e.makeSignUpEndpoint())
	e.r.POST("/signin", e.makeSignInEndpoint())
}
