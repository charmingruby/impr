package memory

import (
	"errors"

	"github.com/charmingruby/impr/service/poll/internal/poll/core/model"
)

type PollSummaryRepository struct {
	Votes   []model.Vote
	Polls   []model.Poll
	Options []model.PollOption
}

func NewPollSummaryRepository() *PollSummaryRepository {
	return &PollSummaryRepository{
		Votes:   []model.Vote{},
		Polls:   []model.Poll{},
		Options: []model.PollOption{},
	}
}

func (r *PollSummaryRepository) FindByPollID(pollID string) (*model.PollSummary, error) {
	var poll *model.Poll
	for _, p := range r.Polls {
		if p.ID == pollID {
			poll = &p
			break
		}
	}

	if poll == nil {
		return nil, errors.New("poll not found")
	}

	optionVotes := make(map[string]int)
	for _, vote := range r.Votes {
		if vote.PollID == pollID {
			optionVotes[vote.PollOptionID]++
		}
	}

	var options []model.PollSummaryOption
	for _, option := range r.Options {
		if option.PollID == pollID {
			options = append(options, model.PollSummaryOption{
				PollOptionID: option.ID,
				Content:      option.Content,
				Votes:        optionVotes[option.ID],
			})
		}
	}

	return &model.PollSummary{
		PollID:    poll.ID,
		Options:   options,
		ExpiresAt: poll.ExpiresAt,
	}, nil
}
