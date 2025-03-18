package repository

import "github.com/charmingruby/impr/service/poll/internal/poll/core/model"

type PollSummaryRepository interface {
	FindByPollID(pollID string) (*model.PollSummary, error)
}
