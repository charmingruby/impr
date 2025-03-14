package service

import (
	"github.com/charmingruby/impr/lib/pkg/core/core_err"
	"github.com/charmingruby/impr/service/poll/internal/poll/core/model"
	"github.com/charmingruby/impr/service/poll/internal/shared/custom_err"
	"github.com/charmingruby/impr/service/poll/test/factory"
)

func (s *Suite) Test_Service_ClosePoll() {
	s.Run("should close an open poll", func() {
		poll := factory.MakePoll(model.Poll{
			Status: model.POLL_OPEN_STATUS,
		})

		err := s.pollRepo.Store(&poll)

		s.NoError(err)
		s.Equal(s.pollRepo.Items[0].Status, model.POLL_OPEN_STATUS)

		err = s.svc.ClosePoll(ClosePollParams{
			PollID:  poll.ID,
			OwnerID: poll.OwnerID,
		})

		s.NoError(err)
		s.Equal(s.pollRepo.Items[0].Status, model.POLL_CLOSED_STATUS)
		s.Equal(len(s.publisher.Messages), 1)
	})

	s.Run("should return an error if poll is already closed", func() {
		poll := factory.MakePoll(model.Poll{
			Status: model.POLL_CLOSED_STATUS,
		})

		err := s.pollRepo.Store(&poll)

		s.NoError(err)
		s.Equal(s.pollRepo.Items[0].Status, model.POLL_CLOSED_STATUS)

		err = s.svc.ClosePoll(ClosePollParams{
			PollID:  poll.ID,
			OwnerID: poll.OwnerID,
		})

		s.Error(err)
		s.Equal(custom_err.NewInvalidActionErr("poll is already closed").Error(), err.Error())
		s.Equal(len(s.publisher.Messages), 0)
	})

	s.Run("should return an error if poll doesn't exists", func() {
		err := s.svc.ClosePoll(ClosePollParams{
			PollID:  "poll-id",
			OwnerID: "owner-id",
		})

		s.Error(err)
		s.Equal(core_err.NewResourceNotFoundErr("poll").Error(), err.Error())
		s.Equal(len(s.publisher.Messages), 0)
	})
}
