package service

import (
	"github.com/charmingruby/impr/lib/pkg/core/core_err"
	"github.com/charmingruby/impr/service/poll/internal/poll/core/model"
	"github.com/charmingruby/impr/service/poll/internal/shared/custom_err"
	"github.com/charmingruby/impr/service/poll/test/factory"
)

func (s *Suite) Test_Service_CreatePoll() {
	s.Run("should create a poll and options without errors", func() {
		params := CreatePollParams{
			Title:              "Color decision",
			Question:           "What is your favorite color?",
			OwnerID:            "owner-id",
			ExpirationTimeInMS: 15 * 60 * 1000, // 15 minutes in ms
			Options:            []string{"Red", "Green"},
		}

		pollID, err := s.svc.CreatePoll(params)

		s.NoError(err)
		s.Equal(1, len(s.pollRepo.Items))
		s.Equal(pollID, s.pollRepo.Items[0].ID)
		s.Equal(s.pollRepo.Items[0].Title, params.Title)
		s.Equal(2, len(s.optionRepo.Items))
		s.Equal(s.optionRepo.Items[0].Content, params.Options[0])
		s.Equal(s.optionRepo.Items[1].Content, params.Options[1])
		s.Equal(len(s.publisher.Messages), 1)
	})

	s.Run("should return an error if poll already exists", func() {
		conflictingPoll := factory.MakePoll(model.Poll{})

		err := s.pollRepo.Store(&conflictingPoll)
		s.NoError(err)

		params := CreatePollParams{
			Title:              conflictingPoll.Title,
			Question:           "What is your favorite color?",
			OwnerID:            conflictingPoll.OwnerID,
			ExpirationTimeInMS: 15 * 60 * 1000, // 15 minutes in ms
			Options:            []string{"Red", "Green"},
		}

		_, err = s.svc.CreatePoll(params)

		s.Error(err)
		s.Equal(core_err.NewConflictErr("title").Error(), err.Error())
		s.Equal(len(s.publisher.Messages), 0)
	})

	s.Run("should return an error if option already exists", func() {
		params := CreatePollParams{
			Title:              "Color decision",
			Question:           "What is your favorite color?",
			OwnerID:            "owner-id",
			ExpirationTimeInMS: 15 * 60 * 1000, // 15 minutes in ms
			Options:            []string{"Red", "Red"},
		}

		_, err := s.svc.CreatePoll(params)

		s.Error(err)

		multipleErr := custom_err.NewMultipleProcessErr(
			[]custom_err.ProcessErr{
				{
					Identitifer: "Red",
					Reason:      core_err.NewConflictErr("content").Error(),
				},
			},
		)

		s.Equal(multipleErr.Error(), err.Error())
		s.Equal(len(s.publisher.Messages), 0)
	})
}
