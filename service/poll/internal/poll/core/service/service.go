package service

import "github.com/charmingruby/impr/service/poll/internal/poll/core/repository"

type Service struct {
	pollRepository repository.PollRepository
}

type Input struct {
	PollRepository repository.PollRepository
}

func New(in Input) *Service {
	return &Service{
		pollRepository: in.PollRepository,
	}
}
