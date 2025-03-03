package endpoint

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/charmingruby/bob/internal/example/core/service"
	"github.com/charmingruby/bob/internal/example/transport/rest/dto/response"
	"github.com/charmingruby/bob/internal/example/transport/rest/dto/request"
	"github.com/charmingruby/bob/internal/shared/custom_err/core_err"
	"github.com/charmingruby/bob/internal/shared/transport/rest"
)

func (e *Endpoint) makeGreetingHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req, err := rest.ParseRequest[request.GreetingRequest](r)
		if err != nil {
			rest.BadRequestErrorResponse(w, err.Error())
			return
		}

		res, err := e.service.Greeting(service.GreetingParams{
			Name: req.Name,
		})

		if err != nil {
			var notFoundErr *core_err.ResourceNotFoundErr
			if errors.As(err, &notFoundErr) {
				rest.NotFoundErrorResponse(w, err.Error())
				return
			}

			slog.Error(err.Error())

			rest.InternalServerErrorResponse(w)
			return
		}

		rest.OKResponse(w, "", response.GreetingResponse{
			Greeting: fmt.Sprintf("Long time no see! `%s` was managed.", res.ID),
		})
	}
}
