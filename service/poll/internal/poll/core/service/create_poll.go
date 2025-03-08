package service

import (
	"github.com/charmingruby/impr/lib/pkg/core/core_err"
	"github.com/charmingruby/impr/service/poll/internal/poll/core/model"
	"github.com/charmingruby/impr/service/poll/internal/shared/custom_err"
)

type CreatePollParams struct {
	Title              string
	Question           string
	ExpirationTimeInMS int
	OwnerID            string
	Options            []string
}

func (s *Service) CreatePoll(params CreatePollParams) error {
	pollExists, err := s.pollRepo.FindByTitleAndOwnerID(params.Title, params.OwnerID)

	if err != nil {
		return custom_err.NewPersistenceErr(err, "find by title and owner id", "poll")
	}

	if pollExists != nil {
		return core_err.NewConflictErr("title")
	}

	poll := model.NewPoll(model.NewPollInput{
		Title:          params.Title,
		Question:       params.Question,
		OwnerID:        params.OwnerID,
		ExpirationTime: params.ExpirationTimeInMS,
	})

	if err := s.pollRepo.Store(poll); err != nil {
		return custom_err.NewPersistenceErr(err, "store", "poll")
	}

	var optionsErrs []custom_err.ProcessErr

	for _, option := range params.Options {
		optExists, err := s.optionRepo.FindByContentAndPollID(option, poll.ID)
		if err != nil {
			optionsErrs = append(optionsErrs, custom_err.ProcessErr{
				Identitifer: option,
				Reason:      err.Error(),
			})

			continue
		}

		if optExists != nil {
			optionsErrs = append(optionsErrs, custom_err.ProcessErr{
				Identitifer: option,
				Reason:      core_err.NewConflictErr("content").Error(),
			})

			continue
		}

		opt := model.NewPollOption(model.NewPollOptionInput{
			Content: option,
			PollID:  poll.ID,
		})

		if err := s.optionRepo.Store(opt); err != nil {
			optionsErrs = append(optionsErrs, custom_err.ProcessErr{
				Identitifer: option,
				Reason:      custom_err.NewPersistenceErr(err, "store", "poll option").Error(),
			})

			continue
		}
	}

	if len(optionsErrs) > 0 {
		return custom_err.NewMultipleProcessErr(optionsErrs)
	}

	return nil
}
