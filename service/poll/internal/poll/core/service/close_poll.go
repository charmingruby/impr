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

type ClosePollParams struct {
	PollID  string
	OwnerID string
}

func (s *Service) ClosePoll(params ClosePollParams) error {
	poll, err := s.pollRepo.FindByIDAndOwnerID(params.PollID, params.OwnerID)

	if err != nil {
		return custom_err.NewPersistenceErr(err, "find by id and owner id", "poll")
	}

	if poll == nil {
		return core_err.NewResourceNotFoundErr("poll")
	}

	if poll.Status == model.POLL_CLOSED_STATUS {
		return custom_err.NewInvalidActionErr("poll is already closed")
	}

	poll.Status = model.POLL_CLOSED_STATUS
	poll.UpdatedAt = time.Now()

	if err = s.pollRepo.Save(poll); err != nil {
		return custom_err.NewPersistenceErr(err, "save", "poll")
	}

	msg, err := event.CreateAuditMessage(event.CreateAuditMessageParams{
		Context:      "poll",
		Subject:      "poll closed",
		Content:      fmt.Sprintf("closed %s poll by %s", poll.ID, poll.OwnerID),
		DispatchedAt: time.Now(),
	})
	if err != nil {
		return err
	}

	if err := s.publisher.Publish(messaging.Message{
		Key:   id.New(),
		Value: msg,
	}); err != nil {
		return err
	}

	return nil
}
