package endpoint

import (
	"errors"

	"github.com/charmingruby/impr/lib/pkg/core_err"
	"github.com/charmingruby/impr/lib/pkg/server/rest"
	"github.com/charmingruby/impr/lib/pkg/validation"
	"github.com/charmingruby/impr/service/identity/internal/account/core/service"
	"github.com/charmingruby/impr/service/identity/internal/shared/custom_err"
	"github.com/charmingruby/impr/service/identity/pkg/logger"
	"github.com/labstack/echo/v4"
)

type ResetPasswordRequest struct {
	Email            string `json:"email" validate:"required,email"`
	ConfirmationCode string `json:"confirmation_code" validate:"required"`
	NewPassword      string `json:"new_password" validate:"required"`
}

func (e *Endpoint) makeResetPasswordEndpoint() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(ResetPasswordRequest)

		if err := c.Bind(req); err != nil {
			return rest.NewPayloadErrorResponse(c, err.Error())
		}

		if err := validation.ValidateStructByTags(req); err != nil {
			return rest.NewPayloadErrorResponse(c, err.Error())
		}

		if err := e.service.ResetPassword(service.ResetPasswordParams{
			Email:            req.Email,
			ConfirmationCode: req.ConfirmationCode,
			NewPassword:      req.NewPassword,
		}); err != nil {
			var mismatchCodeErr *custom_err.InvalidCodeErr
			if errors.As(err, &mismatchCodeErr) {
				return rest.NewUnauthorizedErrorResponse(c, mismatchCodeErr.Error())
			}

			var expiredCodeErr *custom_err.ExpiredCodeErr
			if errors.As(err, &expiredCodeErr) {
				return rest.NewUnauthorizedErrorResponse(c, expiredCodeErr.Error())
			}

			var userNotFoundErr *core_err.ResourceNotFoundErr
			if errors.As(err, &userNotFoundErr) {
				return rest.NewResourceNotFoundErrResponse(c, userNotFoundErr.Error())
			}

			var userNotConfirmedErr *custom_err.UserNotConfirmedErr
			if errors.As(err, &userNotConfirmedErr) {
				return rest.NewUnauthorizedErrorResponse(c, userNotConfirmedErr.Error())
			}

			logger.Log.Error(err.Error())

			return rest.NewInternalServerErrorReponse(c)
		}

		res := map[string]string{
			"message": "password reset successfully",
		}

		return rest.NewOkResponse(c, res)
	}
}
