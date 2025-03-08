package factory

import (
	"time"

	"github.com/charmingruby/impr/lib/pkg/core/id"
	"github.com/charmingruby/impr/service/poll/internal/poll/core/model"
	"github.com/charmingruby/impr/service/poll/pkg/helper"
)

func MakeVote(override model.Vote) model.Vote {
	return model.Vote{
		ID:           helper.Ternary(override.ID == "", id.New(), override.ID),
		PollID:       helper.Ternary(override.PollID == "", id.New(), override.PollID),
		PollOptionID: helper.Ternary(override.PollOptionID == "", id.New(), override.PollOptionID),
		UserID:       helper.Ternary(override.UserID == "", id.New(), override.UserID),
		CreatedAt:    helper.Ternary(override.CreatedAt.UTC().IsZero(), time.Now(), override.CreatedAt),
	}
}
