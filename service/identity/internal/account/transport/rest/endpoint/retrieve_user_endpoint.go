package endpoint

import (
	"errors"

	"github.com/charmingruby/impr/lib/pkg/core_err"
	"github.com/charmingruby/impr/lib/pkg/server/rest"
	"github.com/charmingruby/impr/service/identity/internal/account/core/model"
	"github.com/charmingruby/impr/service/identity/internal/account/core/service"
	"github.com/charmingruby/impr/service/identity/pkg/logger"
	"github.com/labstack/echo/v4"
)

type RetrieveUserData struct {
	User model.User `json:"user"`
}

func (e *Endpoint) makeRetrieveUserEndpoint() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := c.Param("user-id")
		if userID == "" {
			return rest.NewBadRequestResponse(c, "user-id is required")
		}

		user, err := e.service.RetrieveUser(service.RetrieveUserParams{
			ID: userID,
		})
		if err != nil {
			var resourceNotFoundErr *core_err.ResourceNotFoundErr
			if errors.As(err, &resourceNotFoundErr) {
				return rest.NewResourceNotFoundErrResponse(c)
			}

			logger.Log.Error(err.Error())

			return rest.NewInternalServerErrorReponse(c)
		}

		res := map[string]RetrieveUserData{
			"data": {
				User: user,
			},
		}

		return rest.NewOkResponse(c, res)
	}
}
