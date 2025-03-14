package poll

import (
	"github.com/charmingruby/impr/lib/pkg/messaging"
	"github.com/charmingruby/impr/service/poll/internal/poll/core/repository"
	"github.com/charmingruby/impr/service/poll/internal/poll/core/service"
	"github.com/charmingruby/impr/service/poll/internal/poll/database/postgres"
	"github.com/charmingruby/impr/service/poll/internal/poll/transport/rest/endpoint"
	"github.com/charmingruby/impr/service/poll/internal/shared/transport/rest/middleware"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func NewPollRepository(db *sqlx.DB) (repository.PollRepository, error) {
	return postgres.NewPollRepository(db)
}

func NewPollOptionRepository(db *sqlx.DB) (repository.PollOptionRepository, error) {
	return postgres.NewPollOptionRepository(db)
}

func NewVoteRepository(db *sqlx.DB) (repository.VoteRepository, error) {
	return postgres.NewVoteRepository(db)
}

func NewService(
	pollRepo repository.PollRepository,
	optionRepo repository.PollOptionRepository,
	voteRepo repository.VoteRepository,
	publisher messaging.Publisher,
) *service.Service {
	return service.New(pollRepo, optionRepo, voteRepo, publisher)
}

func NewRestHandler(r *echo.Echo, svc *service.Service, mw *middleware.Auth) *endpoint.Endpoint {
	return endpoint.New(r, svc, mw)
}
