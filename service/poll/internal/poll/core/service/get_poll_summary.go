package service

import (
	"github.com/charmingruby/impr/lib/pkg/core/core_err"
	"github.com/charmingruby/impr/service/poll/internal/poll/core/model"
	"github.com/charmingruby/impr/service/poll/internal/shared/custom_err"
)

func (s *Service) GetPollSummary(pollID string) (*model.PollSummary, error) {
	poll, err := s.pollRepo.FindByID(pollID)

	if err != nil {
		return nil, custom_err.NewPersistenceErr(err, "find by id", "poll")
	}

	if poll == nil {
		return nil, core_err.NewResourceNotFoundErr("poll")
	}

	summary, err := s.summaryRepo.FindByPollID(pollID)

	if err != nil {
		return nil, custom_err.NewPersistenceErr(err, "find by poll id", "poll summary")
	}

	if summary == nil {
		return nil, core_err.NewResourceNotFoundErr("poll summary")
	}

	return summary, nil
}
