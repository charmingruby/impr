package repository

import "github.com/charmingruby/impr/service/poll/internal/poll/core/model"

type PollOptionRepository interface {
	FindAllByPollID(pollID string) ([]model.PollOption, error)
	FindByContentAndPollID(content, pollID string) (*model.PollOption, error)
	Store(model *model.PollOption) error
}
