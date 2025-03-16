package factory

import (
	"time"

	"github.com/charmingruby/impr/lib/pkg/core/id"
	"github.com/charmingruby/impr/service/poll/internal/poll/core/model"
	"github.com/charmingruby/impr/service/poll/pkg/helper"
)

func MakePoll(override model.Poll) model.Poll {
	defaultExpiresAt := time.Now().Add(5 * time.Minute)

	return model.Poll{
		ID:        helper.Ternary(override.ID == "", id.New(), override.ID),
		Title:     helper.Ternary(override.Title == "", "Color decision", override.Title),
		Question:  helper.Ternary(override.Question == "", "What is your favorite color?", override.Question),
		Status:    helper.Ternary(override.Status == "", string(model.POLL_OPEN_STATUS), override.Status),
		ExpiresAt: helper.Ternary(override.ExpiresAt == nil, &defaultExpiresAt, override.ExpiresAt),
		OwnerID:   helper.Ternary(override.OwnerID == "", id.New(), override.OwnerID),
		CreatedAt: helper.Ternary(override.CreatedAt.UTC().IsZero(), time.Now(), override.CreatedAt),
		UpdatedAt: helper.Ternary(override.UpdatedAt.UTC().IsZero(), time.Now(), override.UpdatedAt),
	}
}
