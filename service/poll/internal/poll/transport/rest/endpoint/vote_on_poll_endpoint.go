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

type VoteOnPollRequest struct {
	OptionID string `json:"option_id" validate:"required"`
}

func (e *Endpoint) makeVoteOnPollEndpoint() echo.HandlerFunc {
	return func(c echo.Context) error {
		pollID := c.Param("poll_id")

		if pollID == "" {
			return rest.NewBadRequestResponse(c, "poll_id is required")
		}

		hardCodedSampleUserID := "sample-id"

		req := new(VoteOnPollRequest)

		if err := c.Bind(req); err != nil {
			return rest.NewPayloadErrorResponse(c, err.Error())
		}

		if err := validation.ValidateStructByTags(req); err != nil {
			return rest.NewPayloadErrorResponse(c, err.Error())
		}

		voteID, err := e.service.VoteOnPoll(service.VoteOnPollParams{
			PollID:       pollID,
			PollOptionID: req.OptionID,
			UserID:       hardCodedSampleUserID,
		})

		if err != nil {
			var resourceNotFoundErr *core_err.ResourceNotFoundErr
			if errors.As(err, &resourceNotFoundErr) {
				return rest.NewResourceNotFoundErrResponse(c, resourceNotFoundErr.Error())
			}

			var invalidActionErr *custom_err.InvalidActionErr
			if errors.As(err, &invalidActionErr) {
				return rest.NewConflictErrorResponse(c, invalidActionErr.Error())
			}

			logger.Log.Error(err.Error())

			return rest.NewInternalServerErrorReponse(c)
		}

		return rest.NewCreatedResponse(c, "vote", voteID)
	}
}
