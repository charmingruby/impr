package service

import (
	"github.com/charmingruby/impr/lib/pkg/core/core_err"
	"github.com/charmingruby/impr/service/poll/internal/poll/core/model"
	"github.com/charmingruby/impr/service/poll/test/factory"
)

func (s *Suite) Test_Service_GetPollDetails() {
	s.Run("should return poll with options when there is no error", func() {
		poll := factory.MakePoll(model.Poll{})

		err := s.pollRepo.Store(&poll)
		s.NoError(err)

		opt1 := factory.MakePollOption(model.PollOption{
			Content: "option 1",
			PollID:  poll.ID,
		})

		opt2 := factory.MakePollOption(model.PollOption{
			Content: "option 2",
			PollID:  poll.ID,
		})

		err = s.optionRepo.Store(&opt1)
		s.NoError(err)

		err = s.optionRepo.Store(&opt2)
		s.NoError(err)

		result, err := s.svc.GetPollDetails(GetPollDetailsParams{
			PollID: poll.ID,
		})

		s.NoError(err)
		s.Equal(poll.ID, result.Poll.ID)
		s.Equal(2, len(result.Options))
		s.Equal(opt1.ID, result.Options[0].ID)
		s.Equal(opt2.ID, result.Options[1].ID)
	})

	s.Run("should return an error if poll does not exists", func() {
		_, err := s.svc.GetPollDetails(GetPollDetailsParams{
			PollID: "invalid-poll-id",
		})

		s.Error(err)
		s.Equal(core_err.NewResourceNotFoundErr("poll").Error(), err.Error())
	})

	s.Run("should return an empty opts result if there is no poll options", func() {
		poll := factory.MakePoll(model.Poll{})

		err := s.pollRepo.Store(&poll)
		s.NoError(err)

		result, err := s.svc.GetPollDetails(GetPollDetailsParams{
			PollID: poll.ID,
		})

		s.NoError(err)
		s.Equal(poll.ID, result.Poll.ID)
		s.Equal(0, len(result.Options))
	})
}
