package service

import (
	"github.com/charmingruby/impr/service/identity/internal/account/core/model"
)

type FindUserParams struct {
	ID string
}

func (s *Service) FindUser(in FindUserParams) (model.User, error) {
	return s.identityProvider.FindUserByID(in.ID)
}
