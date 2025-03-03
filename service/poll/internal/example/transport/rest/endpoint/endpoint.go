package endpoint

import (
	"github.com/go-chi/chi/v5"
	"github.com/charmingruby/bob/internal/example/core/service"
)

type Endpoint struct {
	router  *chi.Mux
	service *service.Service
}

func New(r *chi.Mux, service *service.Service) *Endpoint {
	return &Endpoint{
		router:  r,
		service: service,
	}
}

func (e *Endpoint) Register() {
	e.router.Post("/example/greeting", e.makeGreetingHandler())
}
