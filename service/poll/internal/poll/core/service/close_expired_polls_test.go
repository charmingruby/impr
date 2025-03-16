package service

import (
	"time"

	"github.com/charmingruby/impr/service/poll/internal/poll/core/model"
	"github.com/charmingruby/impr/service/poll/test/factory"
)

func (s *Suite) Test_Service_CloseExpiredPolls() {
	s.Run("should close expired polls", func() {
		now := time.Now()

		poll1ExpiresAt := now.Add(-time.Hour)
		poll1 := factory.MakePoll(model.Poll{
			ExpiresAt: &poll1ExpiresAt,
			Status:    model.POLL_OPEN_STATUS,
		})

		poll2ExpiresAt := now.Add(-2 * time.Hour)
		poll2 := factory.MakePoll(model.Poll{
			ExpiresAt: &poll2ExpiresAt,
			Status:    model.POLL_OPEN_STATUS,
		})

		err := s.pollRepo.Store(&poll1)
		s.NoError(err)

		err = s.pollRepo.Store(&poll2)
		s.NoError(err)

		err = s.svc.CloseExpiredPolls()
		s.NoError(err)

		closedPoll1, err := s.pollRepo.FindByID(poll1.ID)
		s.NoError(err)
		s.Equal(model.POLL_CLOSED_STATUS, closedPoll1.Status)

		closedPoll2, err := s.pollRepo.FindByID(poll2.ID)
		s.NoError(err)
		s.Equal(model.POLL_CLOSED_STATUS, closedPoll2.Status)
	})

	s.Run("should not close polls that are not expired", func() {
		now := time.Now()

		expiresAt := now.Add(time.Hour)

		poll := factory.MakePoll(model.Poll{
			ExpiresAt: &expiresAt,
			Status:    model.POLL_OPEN_STATUS,
		})

		err := s.pollRepo.Store(&poll)
		s.NoError(err)

		err = s.svc.CloseExpiredPolls()
		s.NoError(err)

		openPoll, err := s.pollRepo.FindByID(poll.ID)
		s.NoError(err)
		s.Equal(model.POLL_OPEN_STATUS, openPoll.Status)
	})

	s.Run("should return error if repository fails", func() {
		s.pollRepo.IsHealthy = false

		err := s.svc.CloseExpiredPolls()
		s.Error(err)
	})
}
