package service

import "github.com/charmingruby/impr/service/poll/internal/poll/core/repository"

type Service struct {
	pollRepository       repository.PollRepository
	pollOptionRepository repository.PollOptionRepository
}

type Input struct {
	PollRepository       repository.PollRepository
	PollOptionRepository repository.PollOptionRepository
}

func New(
	pollRepository repository.PollRepository,
	pollOptionRepository repository.PollOptionRepository,
) *Service {
	return &Service{
		pollRepository:       pollRepository,
		pollOptionRepository: pollOptionRepository,
	}
}
