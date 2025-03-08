package service

import (
	"github.com/charmingruby/impr/lib/pkg/core/core_err"
	"github.com/charmingruby/impr/service/poll/internal/poll/core/model"
	"github.com/charmingruby/impr/service/poll/internal/shared/custom_err"
)

type GetPollDetailsParams struct {
	PollID string
}

type GetPollDetailsResult struct {
	Poll    model.Poll
	Options []model.PollOption
}

func (s *Service) GetPollDetails(params GetPollDetailsParams) (GetPollDetailsResult, error) {
	poll, err := s.pollRepo.FindByID(params.PollID)

	if err != nil {
		return GetPollDetailsResult{}, custom_err.NewPersistenceErr(err, "find by id", "poll")
	}

	if poll == nil {
		return GetPollDetailsResult{}, core_err.NewResourceNotFoundErr("poll")
	}

	options, err := s.optionRepo.FindAllByPollID(params.PollID)

	if err != nil {
		return GetPollDetailsResult{}, custom_err.NewPersistenceErr(err, "find all by poll id", "poll option")
	}

	return GetPollDetailsResult{
		Poll:    *poll,
		Options: options,
	}, nil
}
