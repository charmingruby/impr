package service

import "github.com/charmingruby/impr/service/poll/internal/poll/core/repository"

type Service struct {
	pollRepo   repository.PollRepository
	optionRepo repository.PollOptionRepository
	voteRepo   repository.VoteRepository
}

func New(
	pollRepo repository.PollRepository,
	optionRepo repository.PollOptionRepository,
	voteRepo repository.VoteRepository,
) *Service {
	return &Service{
		pollRepo:   pollRepo,
		optionRepo: optionRepo,
		voteRepo:   voteRepo,
	}
}
