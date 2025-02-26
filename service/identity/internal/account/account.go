package account

import (
	"github.com/charmingruby/impr/service/identity/internal/account/core/gateway"
	"github.com/charmingruby/impr/service/identity/internal/account/core/service"
	"github.com/charmingruby/impr/service/identity/internal/account/transport/rest/endpoint"
	"github.com/labstack/echo/v4"
)

func NewService(identityProvider gateway.IdentityProvider) *service.Service {
	return service.New(identityProvider)
}

func NewRestHandler(r *echo.Echo, svc *service.Service) *endpoint.Endpoint {
	return endpoint.New(r, svc)
}
