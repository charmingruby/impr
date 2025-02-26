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
	api := e.r.Group("/api")

	auth := api.Group("/auth")
	auth.POST("/signup", e.makeSignUpEndpoint())
	auth.POST("/confirm-account", e.makeConfirmAccountEndpoint())
	auth.POST("/signin", e.makeSignInEndpoint())
	auth.POST("/refresh", e.makeRefreshSessionEndpoint())
	auth.POST("/forgot-password", e.makeForgotPasswordEndpoint())
	auth.POST("/reset-password", e.makeResetPasswordEndpoint())

	user := api.Group("/user")
	user.GET("/:user-id", e.makeFindUserEndpoint())
}
