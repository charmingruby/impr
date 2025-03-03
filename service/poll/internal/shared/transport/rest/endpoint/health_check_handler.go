package endpoint

import (
	"net/http"

	"github.com/charmingruby/bob/internal/shared/transport/rest"
)

func (e *Endpoint) makeHealthCheckHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rest.OKResponse(w, "", nil)
	}
}
