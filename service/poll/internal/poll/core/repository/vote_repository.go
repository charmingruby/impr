package repository

import "github.com/charmingruby/impr/service/poll/internal/poll/core/model"

type VoteRepository interface {
	FindByPollIDAndUserID(pollID, userID string) (*model.Vote, error)
	Store(model *model.Vote) error
}
