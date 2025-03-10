package endpoint

import (
	"errors"

	"github.com/charmingruby/impr/lib/pkg/core/core_err"
	"github.com/charmingruby/impr/lib/pkg/http/server/rest"
	"github.com/charmingruby/impr/lib/pkg/validation"
	"github.com/charmingruby/impr/service/poll/internal/poll/core/service"
	"github.com/charmingruby/impr/service/poll/internal/shared/custom_err"
	"github.com/charmingruby/impr/service/poll/pkg/logger"
	"github.com/labstack/echo/v4"
)

type CreatePollRequest struct {
	Title              string   `json:"title" validate:"required,min=3,max=144"`
	Question           string   `json:"question" validate:"required,min=3,max=144"`
	ExpirationTimeInMS int      `json:"expiration_time_in_ms" validate:"required,min=1"`
	OwnerID            string   `json:"owner_id" validate:"required"`
	Options            []string `json:"options" validate:"required,min=2"`
}

func (e *Endpoint) makeCreatePollEndpoint() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(CreatePollRequest)

		if err := c.Bind(req); err != nil {
			return rest.NewPayloadErrorResponse(c, err.Error())
		}

		if err := validation.ValidateStructByTags(req); err != nil {
			return rest.NewPayloadErrorResponse(c, err.Error())
		}

		pollID, err := e.service.CreatePoll(service.CreatePollParams{
			Title:              req.Title,
			Question:           req.Question,
			ExpirationTimeInMS: req.ExpirationTimeInMS,
			OwnerID:            req.OwnerID,
			Options:            req.Options,
		})

		if err != nil {
			var multipleProcessErr *custom_err.MultipleProcessErr
			if errors.As(err, &multipleProcessErr) {
				return rest.NewUnprocessableEntity(c, multipleProcessErr.Error())
			}

			var conflictErr *core_err.ConflictErr
			if errors.As(err, &conflictErr) {
				return rest.NewConflictErrorResponse(c, conflictErr.Error())
			}

			logger.Log.Error(err.Error())

			return rest.NewInternalServerErrorReponse(c)
		}

		return rest.NewCreatedResponse(c, "poll", pollID)
	}
}
