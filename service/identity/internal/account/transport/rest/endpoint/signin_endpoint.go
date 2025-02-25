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

type SignInRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type SignInData struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
}

func (e *Endpoint) makeSignInEndpoint() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(SignInRequest)

		if err := c.Bind(req); err != nil {
			return rest.NewPayloadErrorResponse(c, err.Error())
		}

		if err := validation.ValidateStructByTags(req); err != nil {
			return rest.NewPayloadErrorResponse(c, err.Error())
		}

		tokens, err := e.service.SignIn(service.SignInParams{
			Email:    req.Email,
			Password: req.Password,
		})

		if err != nil {
			var invalidCredentialsErr *custom_err.InvalidCredentialsErr
			if errors.As(err, &invalidCredentialsErr) {
				return rest.NewUnauthorizedErrorResponse(c, invalidCredentialsErr.Error())
			}

			var userNotConfirmedErr *custom_err.UserNotConfirmedErr
			if errors.As(err, &userNotConfirmedErr) {
				return rest.NewBadRequestResponse(c, userNotConfirmedErr.Error())
			}

			logger.Log.Error(err.Error())

			return rest.NewInternalServerErrorReponse(c)
		}

		res := map[string]SignInData{
			"data": {
				RefreshToken: tokens.RefreshToken,
				AccessToken:  tokens.AccessToken,
			},
		}

		return rest.NewOkResponse(c, res)
	}
}
