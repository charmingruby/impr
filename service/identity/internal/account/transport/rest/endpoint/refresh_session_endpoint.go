package endpoint

import (
	"errors"

	"github.com/charmingruby/impr/lib/pkg/server/rest"
	"github.com/charmingruby/impr/lib/pkg/validation"
	"github.com/charmingruby/impr/service/identity/internal/account/core/service"
	"github.com/charmingruby/impr/service/identity/internal/shared/custom_err"
	"github.com/charmingruby/impr/service/identity/pkg/logger"
	"github.com/labstack/echo/v4"
)

type RefreshSessionRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type RefreshSessionData struct {
	AccessToken string `json:"access_token"`
}

func (e *Endpoint) makeRefreshSessionEndpoint() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(RefreshSessionRequest)

		if err := c.Bind(req); err != nil {
			return rest.NewPayloadErrorResponse(c, err.Error())
		}

		if err := validation.ValidateStructByTags(req); err != nil {
			return rest.NewPayloadErrorResponse(c, err.Error())
		}

		token, err := e.service.RefreshSession(service.RefreshSessionParams{
			RefreshToken: req.RefreshToken,
		})
		if err != nil {
			var invalidTokenErr *custom_err.InvalidTokenErr
			if errors.As(err, &invalidTokenErr) {
				return rest.NewUnauthorizedErrorResponse(c, invalidTokenErr.Error())
			}

			logger.Log.Error(err.Error())

			return rest.NewInternalServerErrorReponse(c)
		}

		res := map[string]RefreshSessionData{
			"data": {
				AccessToken: token.RenewedAccessToken,
			},
		}

		return rest.NewOkResponse(c, res)
	}
}
