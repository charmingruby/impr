package service

import (
	"fmt"
	"time"

	"github.com/charmingruby/impr/service/poll/internal/poll/core/model"
	"github.com/charmingruby/impr/service/poll/internal/shared/custom_err"
	"github.com/charmingruby/impr/service/poll/pkg/logger"
)

func (s *Service) CloseExpiredPolls() error {
	now := time.Now()

	pollsToBeClosed, err := s.pollRepo.FindAllWithInferiorExpiresAt(now)

	if err != nil {
		return custom_err.NewPersistenceErr(err, "find all with inferior expires at", "poll")
	}

	if len(pollsToBeClosed) == 0 {
		return nil
	}

	var errs []custom_err.ProcessErr

	for _, poll := range pollsToBeClosed {
		logger.Log.Debug(fmt.Sprintf("Closing poll: %s", poll.ID))

		poll.Status = model.POLL_CLOSED_STATUS

		if err := s.pollRepo.Save(&poll); err != nil {
			errs = append(errs, custom_err.ProcessErr{
				Identitifer: poll.ID,
				Reason:      err.Error(),
			})
		}
	}

	if len(errs) != 0 {
		return custom_err.NewMultipleProcessErr(errs)
	}

	return nil
}
