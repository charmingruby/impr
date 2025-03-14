package service

import (
	"github.com/charmingruby/impr/lib/pkg/messaging"
	"github.com/charmingruby/impr/service/poll/internal/poll/core/repository"
)

type Service struct {
	pollRepo   repository.PollRepository
	optionRepo repository.PollOptionRepository
	voteRepo   repository.VoteRepository
	publisher  messaging.Publisher
}

func New(
	pollRepo repository.PollRepository,
	optionRepo repository.PollOptionRepository,
	voteRepo repository.VoteRepository,
	publisher messaging.Publisher,
) *Service {
	return &Service{
		pollRepo:   pollRepo,
		optionRepo: optionRepo,
		voteRepo:   voteRepo,
		publisher:  publisher,
	}
}
