package endpoint

import (
	"errors"

	"github.com/charmingruby/impr/lib/pkg/core_err"
	"github.com/charmingruby/impr/lib/pkg/server/rest"
	"github.com/charmingruby/impr/lib/pkg/validation"
	"github.com/charmingruby/impr/service/identity/internal/account/core/service"
	"github.com/charmingruby/impr/service/identity/pkg/logger"
	"github.com/labstack/echo/v4"
)

type ForgotPasswordRequest struct {
	Email string `json:"email" validate:"required,email"`
}

func (e *Endpoint) makeForgotPasswordEndpoint() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(ForgotPasswordRequest)

		if err := c.Bind(req); err != nil {
			return rest.NewPayloadErrorResponse(c, err.Error())
		}

		if err := validation.ValidateStructByTags(req); err != nil {
			return rest.NewPayloadErrorResponse(c, err.Error())
		}

		if err := e.service.ForgotPassword(service.ForgotPasswordParams{
			Email: req.Email,
		}); err != nil {
			var resourceNotFoundErr *core_err.ResourceNotFoundErr
			if errors.As(err, &resourceNotFoundErr) {
				return rest.NewResourceNotFoundErrResponse(c, resourceNotFoundErr.Error())
			}

			logger.Log.Error(err.Error())

			return rest.NewInternalServerErrorReponse(c)
		}

		res := map[string]string{
			"message": "reset code sent to email",
		}

		return rest.NewOkResponse(c, res)
	}
}
