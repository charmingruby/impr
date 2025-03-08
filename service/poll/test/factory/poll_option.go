package factory

import (
	"time"

	"github.com/charmingruby/impr/lib/pkg/core/id"
	"github.com/charmingruby/impr/service/poll/internal/poll/core/model"
	"github.com/charmingruby/impr/service/poll/pkg/helper"
)

func MakePollOption(override model.PollOption) model.PollOption {
	return model.PollOption{
		ID:        helper.Ternary(override.ID == "", id.New(), override.ID),
		Content:   helper.Ternary(override.Content == "", "Red", override.Content),
		PollID:    helper.Ternary(override.PollID == "", id.New(), override.PollID),
		CreatedAt: helper.Ternary(override.CreatedAt.UTC().IsZero(), time.Now(), override.CreatedAt),
	}
}
