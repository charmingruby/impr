package service

import (
	"fmt"
	"time"

	"github.com/charmingruby/impr/lib/pkg/core/core_err"
	"github.com/charmingruby/impr/lib/pkg/core/id"
	"github.com/charmingruby/impr/lib/pkg/messaging"
	"github.com/charmingruby/impr/service/poll/internal/poll/core/event"
	"github.com/charmingruby/impr/service/poll/internal/poll/core/model"
	"github.com/charmingruby/impr/service/poll/internal/shared/custom_err"
)

type VoteOnPollParams struct {
	PollID       string
	PollOptionID string
	UserID       string
}

func (s *Service) VoteOnPoll(params VoteOnPollParams) (string, error) {
	poll, err := s.pollRepo.FindByID(params.PollID)

	if err != nil {
		return "", custom_err.NewPersistenceErr(err, "find by id", "poll")
	}

	if poll == nil {
		return "", core_err.NewResourceNotFoundErr("poll")
	}

	if poll.Status != model.POLL_OPEN_STATUS {
		return "", custom_err.NewInvalidActionErr("poll is not open")
	}

	voteExists, err := s.voteRepo.FindByPollIDAndUserID(params.PollID, params.UserID)

	if err != nil {
		return "", custom_err.NewPersistenceErr(err, "find vote by poll id and user id", "vote")
	}

	if voteExists != nil {
		return "", custom_err.NewInvalidActionErr("vote already exists")
	}

	option, err := s.optionRepo.FindByID(params.PollOptionID)

	if err != nil {
		return "", custom_err.NewPersistenceErr(err, "find by id", "poll option")
	}

	if option == nil {
		return "", core_err.NewResourceNotFoundErr("poll option")
	}

	vote := model.NewVote(model.NewVoteInput{
		PollID:       params.PollID,
		UserID:       params.UserID,
		PollOptionID: params.PollOptionID,
	})

	if err := s.voteRepo.Store(vote); err != nil {
		return "", custom_err.NewPersistenceErr(err, "store", "vote")
	}

	msg, err := event.CreateAuditMessage(event.CreateAuditMessageParams{
		Context:      "vote",
		Subject:      "vote submited",
		Content:      fmt.Sprintf("%s submited vote: %s, on poll: %s", poll.OwnerID, params.PollOptionID, poll.ID),
		DispatchedAt: time.Now(),
	})
	if err != nil {
		return "", err
	}

	if err := s.publisher.Publish(messaging.Message{
		Key:   id.New(),
		Value: msg,
	}); err != nil {
		return "", err
	}

	return vote.ID, nil
}
