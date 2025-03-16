package endpoint

import (
	"errors"
	"fmt"
	"time"

	"github.com/charmingruby/impr/lib/pkg/core/core_err"
	"github.com/charmingruby/impr/lib/pkg/http/server/rest"
	"github.com/charmingruby/impr/lib/pkg/validation"
	"github.com/charmingruby/impr/service/poll/internal/poll/core/service"
	"github.com/charmingruby/impr/service/poll/internal/shared/custom_err"
	"github.com/charmingruby/impr/service/poll/internal/shared/transport/rest/middleware"
	"github.com/charmingruby/impr/service/poll/pkg/logger"
	"github.com/labstack/echo/v4"
)

type CreatePollRequest struct {
	Title     string   `json:"title" validate:"required,min=3,max=144"`
	Question  string   `json:"question" validate:"required,min=3,max=144"`
	ExpiresAt string   `json:"expires_at" `
	Options   []string `json:"options" validate:"required,min=2"`
}

func (e *Endpoint) makeCreatePollEndpoint() echo.HandlerFunc {
	return func(c echo.Context) error {
		accountID := fmt.Sprintf("%v", c.Get(middleware.ACCOUNT_ID_KEY))
		if accountID == "" {
			return rest.NewUnauthorizedErrorResponse(c, "invalid or expired token")
		}

		req := new(CreatePollRequest)

		if err := c.Bind(req); err != nil {
			return rest.NewPayloadErrorResponse(c, err.Error())
		}

		if err := validation.ValidateStructByTags(req); err != nil {
			return rest.NewPayloadErrorResponse(c, err.Error())
		}

		var expiresAt *time.Time
		if req.ExpiresAt != "" {
			parsedTime, err := time.Parse(time.RFC3339, req.ExpiresAt)
			if err != nil {
				return rest.NewPayloadErrorResponse(c, "invalid expires_at format, expected RFC3339")
			}

			expiresAt = &parsedTime
		}

		pollID, err := e.service.CreatePoll(service.CreatePollParams{
			Title:     req.Title,
			Question:  req.Question,
			ExpiresAt: expiresAt,
			OwnerID:   accountID,
			Options:   req.Options,
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
