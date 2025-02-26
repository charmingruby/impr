package service

import (
	"github.com/charmingruby/impr/service/identity/internal/account/core/model"
)

type RetrieveUserParams struct {
	ID string
}

func (s *Service) RetrieveUser(in RetrieveUserParams) (model.User, error) {
	return s.identityProvider.RetrieveUser(in.ID)
}
