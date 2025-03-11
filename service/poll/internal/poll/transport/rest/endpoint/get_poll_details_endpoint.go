package endpoint

import (
	"errors"

	"github.com/charmingruby/impr/lib/pkg/core/core_err"
	"github.com/charmingruby/impr/lib/pkg/http/server/rest"
	"github.com/charmingruby/impr/service/poll/internal/poll/core/service"
	"github.com/charmingruby/impr/service/poll/pkg/logger"
	"github.com/labstack/echo/v4"
)

func (e *Endpoint) makeGetPollDetailsEndpoint() echo.HandlerFunc {
	return func(c echo.Context) error {
		pollID := c.Param("poll_id")

		if pollID == "" {
			return rest.NewBadRequestResponse(c, "poll_id is required")
		}

		result, err := e.service.GetPollDetails(service.GetPollDetailsParams{
			PollID: pollID,
		})

		if err != nil {

			var resourceNotFoundErr *core_err.ResourceNotFoundErr
			if errors.As(err, &resourceNotFoundErr) {
				return rest.NewResourceNotFoundErrResponse(c, resourceNotFoundErr.Error())
			}

			logger.Log.Error(err.Error())

			return rest.NewInternalServerErrorReponse(c)
		}

		return rest.NewOkResponse(c, result)
	}
}
