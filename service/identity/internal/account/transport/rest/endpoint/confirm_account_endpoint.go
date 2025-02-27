package endpoint

import (
	"errors"

	"github.com/charmingruby/impr/lib/pkg/http/server/rest"
	"github.com/charmingruby/impr/lib/pkg/validation"
	"github.com/charmingruby/impr/service/identity/internal/account/core/service"
	"github.com/charmingruby/impr/service/identity/internal/shared/custom_err"
	"github.com/charmingruby/impr/service/identity/pkg/logger"
	"github.com/labstack/echo/v4"
)

type ConfirmAccountRequest struct {
	Email            string `json:"email" validate:"required,email"`
	ConfirmationCode string `json:"confirmation_code" validate:"required"`
}

func (e *Endpoint) makeConfirmAccountEndpoint() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(ConfirmAccountRequest)

		if err := c.Bind(req); err != nil {
			return rest.NewPayloadErrorResponse(c, err.Error())
		}

		if err := validation.ValidateStructByTags(req); err != nil {
			return rest.NewPayloadErrorResponse(c, err.Error())
		}

		if err := e.service.ConfirmAccount(service.ConfirmAccountParams{
			Email: req.Email,
			Code:  req.ConfirmationCode,
		}); err != nil {
			var mismatchCodeErr *custom_err.InvalidCodeErr
			if errors.As(err, &mismatchCodeErr) {
				return rest.NewUnauthorizedErrorResponse(c, mismatchCodeErr.Error())
			}

			var invalidCredentialsErr *custom_err.InvalidCredentialsErr
			if errors.As(err, &invalidCredentialsErr) {
				return rest.NewUnauthorizedErrorResponse(c, invalidCredentialsErr.Error())
			}

			var expiredCodeErr *custom_err.ExpiredCodeErr
			if errors.As(err, &expiredCodeErr) {
				return rest.NewUnauthorizedErrorResponse(c, expiredCodeErr.Error())
			}

			logger.Log.Error(err.Error())

			return rest.NewInternalServerErrorReponse(c)
		}

		res := map[string]string{
			"message": "account confirmed successfully",
		}

		return rest.NewOkResponse(c, res)
	}
}
