package service

import (
	"github.com/charmingruby/impr/service/identity/internal/account/core/model"
	"github.com/charmingruby/impr/service/identity/internal/shared/custom_err/core_err"
)

type GreetingParams struct {
	Name string
}

type GreetingResult struct {
	ID string
}

func (s *Service) Greeting(params GreetingParams) (GreetingResult, error) {
	user := model.NewUser(model.NewUserInput{
		Name: params.Name,
	})

	if err := s.userRepository.Store(user); err != nil {
		return GreetingResult{}, err
	}

	user.SoftDelete()
	if err := s.userRepository.Delete(user); err != nil {
		return GreetingResult{}, err
	}

	userFound, err := s.userRepository.FindByID(user.ID)
	if err != nil {
		return GreetingResult{}, err
	}

	if userFound == nil {
		return GreetingResult{}, core_err.NewResourceNotFoundErr("user")
	}

	return GreetingResult{
		ID: userFound.ID,
	}, nil
}
