package service

import (
	"github.com/charmingruby/bob/internal/example/core/model"
	"github.com/charmingruby/bob/internal/shared/custom_err/core_err"
)

type GreetingParams struct {
	Name string
}

type GreetingResult struct {
	ID string
}

func (s *Service) Greeting(params GreetingParams) (GreetingResult, error) {
	example := model.NewExample(model.NewExampleInput{
		Name: params.Name,
	})
	
	if err := s.exampleRepository.Store(example); err != nil {
		return GreetingResult{}, err
	}

	example.SoftDelete()
	if err := s.exampleRepository.Delete(example); err != nil {
		return GreetingResult{}, err
	}

	exampleFound, err := s.exampleRepository.FindByID(example.ID)
	if err != nil {
		return GreetingResult{}, err
	}
	
	if exampleFound == nil {
		return GreetingResult{}, core_err.NewResourceNotFoundErr("example")
	}

	return GreetingResult{
		ID: exampleFound.ID,
	}, nil
}
