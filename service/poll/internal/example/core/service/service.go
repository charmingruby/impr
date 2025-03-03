package service

import "github.com/charmingruby/bob/internal/example/core/repository"

type Service struct {
    exampleRepository repository.ExampleRepository
}

type Input struct {
    ExampleRepository repository.ExampleRepository
}

func New(in Input) *Service {
	return &Service{
        exampleRepository: in.ExampleRepository,
    }
}
