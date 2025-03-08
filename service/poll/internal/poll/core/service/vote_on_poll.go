package service

import (
	"github.com/charmingruby/impr/lib/pkg/core/core_err"
	"github.com/charmingruby/impr/service/poll/internal/poll/core/model"
	"github.com/charmingruby/impr/service/poll/internal/shared/custom_err"
)

type VoteOnPollParams struct {
	PollID       string
	PollOptionID string
	UserID       string
}

func (s *Service) VoteOnPoll(params VoteOnPollParams) error {
	poll, err := s.pollRepo.FindByID(params.PollID)

	if err != nil {
		return custom_err.NewPersistenceErr(err, "find by id", "poll")
	}

	if poll == nil {
		return core_err.NewResourceNotFoundErr("poll")
	}

	voteExists, err := s.voteRepo.FindByPollIDAndUserID(params.PollOptionID, params.UserID)

	if err != nil {
		return custom_err.NewPersistenceErr(err, "find vote by poll id and user id", "vote")
	}

	if voteExists != nil {
		return custom_err.NewInvalidActionErr("vote already exists")
	}

	option, err := s.optionRepo.FindByID(params.PollOptionID)

	if err != nil {
		return custom_err.NewPersistenceErr(err, "find by id", "poll option")
	}

	if option == nil {
		return core_err.NewResourceNotFoundErr("poll option")
	}

	vote := model.NewVote(model.NewVoteInput{
		PollID:       params.PollID,
		UserID:       params.UserID,
		PollOptionID: params.PollOptionID,
	})

	if err := s.voteRepo.Store(vote); err != nil {
		return custom_err.NewPersistenceErr(err, "store", "vote")
	}

	return nil
}
