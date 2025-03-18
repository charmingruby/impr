package service

import (
	"github.com/charmingruby/impr/lib/pkg/messaging"
	"github.com/charmingruby/impr/service/poll/internal/poll/core/repository"
)

type Service struct {
	pollRepo    repository.PollRepository
	optionRepo  repository.PollOptionRepository
	voteRepo    repository.VoteRepository
	summaryRepo repository.PollSummaryRepository
	publisher   messaging.Publisher
}

func New(
	pollRepo repository.PollRepository,
	optionRepo repository.PollOptionRepository,
	voteRepo repository.VoteRepository,
	summaryRepo repository.PollSummaryRepository,
	publisher messaging.Publisher,
) *Service {
	return &Service{
		pollRepo:    pollRepo,
		optionRepo:  optionRepo,
		voteRepo:    voteRepo,
		summaryRepo: summaryRepo,
		publisher:   publisher,
	}
}
