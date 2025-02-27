package endpoint

import (
	"errors"

	"github.com/charmingruby/impr/lib/pkg/core/core_err"
	"github.com/charmingruby/impr/lib/pkg/http/server/rest"
	"github.com/charmingruby/impr/lib/pkg/validation"
	"github.com/charmingruby/impr/service/identity/internal/account/core/service"
	"github.com/charmingruby/impr/service/identity/pkg/logger"
	"github.com/labstack/echo/v4"
)

type SignUpRequest struct {
	FirstName string `json:"first_name" validate:"required,min=3,max=72"`
	LastName  string `json:"last_name" validate:"required,min=3,max=144"`
	Email     string `json:"email" validate:"required,email"`
	Birthdate string `json:"birthdate" validate:"required"`
	Password  string `json:"password" validate:"required,min=10,max=64"`
}

func (e *Endpoint) makeSignUpEndpoint() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(SignUpRequest)

		if err := c.Bind(req); err != nil {
			return rest.NewPayloadErrorResponse(c, err.Error())
		}

		if err := validation.ValidateStructByTags(req); err != nil {
			return rest.NewPayloadErrorResponse(c, err.Error())
		}

		userID, err := e.service.SignUp(service.SignUpParams{
			FirstName: req.FirstName,
			LastName:  req.LastName,
			Email:     req.Email,
			Birthdate: req.Birthdate,
			Password:  req.Password,
		})
		if err != nil {
			var invalidFieldFormatErr *core_err.InvalidFieldFormatErr
			if errors.As(err, &invalidFieldFormatErr) {
				return rest.NewPayloadErrorResponse(c, invalidFieldFormatErr.Error())
			}

			var conflictErr *core_err.ConflictErr
			if errors.As(err, &conflictErr) {
				return rest.NewConflictErrorResponse(c, conflictErr.Error())
			}

			logger.Log.Error(err.Error())

			return rest.NewInternalServerErrorReponse(c)
		}

		return rest.NewCreatedResponse(c, "user", userID)
	}
}
