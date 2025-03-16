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

type CreatePollParams struct {
	Title     string
	Question  string
	ExpiresAt *time.Time
	OwnerID   string
	Options   []string
}

func (s *Service) CreatePoll(params CreatePollParams) (string, error) {
	pollExists, err := s.pollRepo.FindByTitleAndOwnerID(params.Title, params.OwnerID)

	if err != nil {
		return "", custom_err.NewPersistenceErr(err, "find by title and owner id", "poll")
	}

	if pollExists != nil {
		return "", core_err.NewConflictErr("title")
	}

	poll := model.NewPoll(model.NewPollInput{
		Title:     params.Title,
		Question:  params.Question,
		OwnerID:   params.OwnerID,
		ExpiresAt: params.ExpiresAt,
	})

	if err := s.pollRepo.Store(poll); err != nil {
		return "", custom_err.NewPersistenceErr(err, "store", "poll")
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
		return "", custom_err.NewMultipleProcessErr(optionsErrs)
	}

	msg, err := event.CreateAuditMessage(event.CreateAuditMessageParams{
		Context:      "poll",
		Subject:      "poll created",
		Content:      fmt.Sprintf("%s created new poll: %s", poll.OwnerID, poll.ID),
		DispatchedAt: time.Now(),
	})
	if err != nil {
		return "", err
	}

	if err := s.publisher.Publish(messaging.Message{
		Key:   id.New(),
		Value: msg,
	}); err != nil {
		return "", err
	}

	return poll.ID, nil
}
