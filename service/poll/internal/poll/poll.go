package poll

import (
	"github.com/charmingruby/impr/service/poll/internal/poll/core/service"
	"github.com/charmingruby/impr/service/poll/internal/poll/database/postgres"
	"github.com/charmingruby/impr/service/poll/internal/poll/transport/rest/endpoint"
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func NewService(db *sqlx.DB) (*service.Service, error) {
	pollRepo, err := postgres.NewPollRepository(db)
	if err != nil {
		return nil, err
	}

	return service.New(service.Input{
		PollRepository: pollRepo,
	}), nil
}

func NewHTTPHandler(r *chi.Mux, service *service.Service) {
	endpoint.New(r, service).Register()
}
