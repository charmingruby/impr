package middleware

import (
	"strings"

	"github.com/charmingruby/impr/lib/pkg/http/server/rest"
	"github.com/charmingruby/impr/service/poll/internal/shared/transport/grpc/client"
	"github.com/labstack/echo/v4"
)

const ACCOUNT_ID_KEY = "accountID"

type Auth struct {
	service *client.Service
}

func NewAuth(service *client.Service) *Auth {
	return &Auth{
		service: service,
	}
}

func (a *Auth) Intercept(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return rest.NewUnauthorizedErrorResponse(c, "missing token")
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return rest.NewUnauthorizedErrorResponse(c, "invalid token format")
		}

		token := parts[1]

		res, err := a.service.VerifyToken(token)
		if err != nil || !res.IsValid {
			return rest.NewUnauthorizedErrorResponse(c, "invalid or expired token")
		}

		c.Set(ACCOUNT_ID_KEY, res.AccountID)

		return next(c)
	}
}
