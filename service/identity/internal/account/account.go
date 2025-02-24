package account

import (
	"github.com/charmingruby/impr/lib/pkg/client/service/awsc"
	"github.com/charmingruby/impr/service/identity/internal/account/core/service"
	"github.com/charmingruby/impr/service/identity/internal/account/transport/rest/client"
	"github.com/charmingruby/impr/service/identity/internal/account/transport/rest/endpoint"
	"github.com/labstack/echo/v4"
)

func NewService(cognitoClient *awsc.CognitoClient) *service.Service {
	identityProviderClient := client.NewCognitoIdentityProvider(cognitoClient)

	return service.New(identityProviderClient)
}

func NewRestHandler(r *echo.Echo, svc *service.Service) *endpoint.Endpoint {
	return endpoint.New(r, svc)
}
