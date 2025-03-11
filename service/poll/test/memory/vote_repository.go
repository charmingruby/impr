package memory

import (
	"fmt"

	"github.com/charmingruby/impr/service/poll/internal/poll/core/model"
)

type VoteRepository struct {
	Items     []model.Vote
	IsHealthy bool
}

func NewVoteRepository() *VoteRepository {
	return &VoteRepository{
		Items:     []model.Vote{},
		IsHealthy: true,
	}
}

func (r *VoteRepository) FindByPollIDAndUserID(pollID, userID string) (*model.Vote, error) {
	for _, item := range r.Items {
		if item.PollID == pollID && item.UserID == userID {
			return &item, nil
		}
	}

	if !r.IsHealthy {
		return nil, fmt.Errorf("vote datasource is unhealthy")
	}

	return nil, nil
}

func (r *VoteRepository) Store(model *model.Vote) error {
	if !r.IsHealthy {
		return fmt.Errorf("vote datasource is unhealthy")
	}

	r.Items = append(r.Items, *model)

	return nil
}
