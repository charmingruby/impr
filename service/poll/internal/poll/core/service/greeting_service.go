package service

import (
	"github.com/charmingruby/impr/service/poll/internal/poll/core/model"
	"github.com/charmingruby/impr/service/poll/internal/shared/custom_err/core_err"
)

type GreetingParams struct {
	Name string
}

type GreetingResult struct {
	ID string
}

func (s *Service) Greeting(params GreetingParams) (GreetingResult, error) {
	poll := model.NewPoll(model.NewPollInput{
		Name: params.Name,
	})

	if err := s.pollRepository.Store(poll); err != nil {
		return GreetingResult{}, err
	}

	if err := s.pollRepository.Delete(poll); err != nil {
		return GreetingResult{}, err
	}

	pollFound, err := s.pollRepository.FindByID(poll.ID)
	if err != nil {
		return GreetingResult{}, err
	}

	if pollFound == nil {
		return GreetingResult{}, core_err.NewResourceNotFoundErr("poll")
	}

	return GreetingResult{
		ID: pollFound.ID,
	}, nil
}
