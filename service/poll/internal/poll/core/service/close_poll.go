package service

import (
	"time"

	"github.com/charmingruby/impr/lib/pkg/core/core_err"
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

	return nil
}
