package poll

import (
	"github.com/charmingruby/impr/service/poll/internal/poll/core/repository"
	"github.com/charmingruby/impr/service/poll/internal/poll/core/service"
	"github.com/charmingruby/impr/service/poll/internal/poll/transport/rest/endpoint"
	"github.com/labstack/echo/v4"
)

func NewService(
	pollRepo repository.PollRepository,
	optionRepo repository.PollOptionRepository,
	voteRepo repository.VoteRepository,
) *service.Service {
	return service.New(pollRepo, optionRepo, voteRepo)
}

func NewRestHandler(r *echo.Echo, svc *service.Service) *endpoint.Endpoint {
	return endpoint.New(r, svc)
}
