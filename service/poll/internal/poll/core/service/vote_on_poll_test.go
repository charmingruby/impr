package service

import (
	"github.com/charmingruby/impr/lib/pkg/core/core_err"
	"github.com/charmingruby/impr/service/poll/internal/poll/core/model"
	"github.com/charmingruby/impr/service/poll/internal/shared/custom_err"
	"github.com/charmingruby/impr/service/poll/test/factory"
)

func (s *Suite) Test_Service_VoteOnPoll() {
	s.Run("should return nil when there is no error", func() {
		poll := factory.MakePoll(model.Poll{})

		err := s.pollRepo.Store(&poll)
		s.NoError(err)

		opt := factory.MakePollOption(model.PollOption{
			Content: "option 1",
			PollID:  poll.ID,
		})

		err = s.optionRepo.Store(&opt)
		s.NoError(err)

		voteID, err := s.svc.VoteOnPoll(VoteOnPollParams{
			PollID:       poll.ID,
			PollOptionID: opt.ID,
			UserID:       "valid-user-id",
		})

		s.NoError(err)
		s.Equal(1, len(s.voteRepo.Items))
		s.Equal(voteID, s.voteRepo.Items[0].ID)
		s.Equal("valid-user-id", s.voteRepo.Items[0].UserID)
		s.Equal(poll.ID, s.voteRepo.Items[0].PollID)
		s.Equal(opt.ID, s.voteRepo.Items[0].PollOptionID)
		s.Equal(len(s.publisher.Messages), 1)
	})

	s.Run("should return an error if poll doesn't exists", func() {
		invalidPollID := "invalid-poll-id"

		opt := factory.MakePollOption(model.PollOption{
			Content: "option 1",
			PollID:  invalidPollID,
		})

		err := s.optionRepo.Store(&opt)
		s.NoError(err)

		_, err = s.svc.VoteOnPoll(VoteOnPollParams{
			PollID:       invalidPollID,
			PollOptionID: opt.ID,
			UserID:       "valid-user-id",
		})

		s.Error(err)
		s.Equal(core_err.NewResourceNotFoundErr("poll").Error(), err.Error())
		s.Equal(len(s.publisher.Messages), 0)
	})

	s.Run("should return an error if option doesn't exists", func() {
		poll := factory.MakePoll(model.Poll{})

		err := s.pollRepo.Store(&poll)
		s.NoError(err)

		_, err = s.svc.VoteOnPoll(VoteOnPollParams{
			PollID:       poll.ID,
			PollOptionID: "invalid-option-id",
			UserID:       "valid-user-id",
		})

		s.Error(err)
		s.Equal(core_err.NewResourceNotFoundErr("poll option").Error(), err.Error())
		s.Equal(len(s.publisher.Messages), 0)
	})

	s.Run("should return an error if user already voted on poll", func() {
		poll := factory.MakePoll(model.Poll{})

		err := s.pollRepo.Store(&poll)
		s.NoError(err)

		opt := factory.MakePollOption(model.PollOption{
			Content: "option 1",
			PollID:  poll.ID,
		})

		err = s.optionRepo.Store(&opt)
		s.NoError(err)

		userID := "valid-user-id"

		vote := factory.MakeVote(model.Vote{
			PollID:       poll.ID,
			PollOptionID: opt.ID,
			UserID:       userID,
		})

		err = s.voteRepo.Store(&vote)
		s.NoError(err)

		_, err = s.svc.VoteOnPoll(VoteOnPollParams{
			PollID:       poll.ID,
			PollOptionID: opt.ID,
			UserID:       userID,
		})

		s.Error(err)
		s.Equal(custom_err.NewInvalidActionErr("vote already exists").Error(), err.Error())
		s.Equal(len(s.publisher.Messages), 0)
	})

	s.Run("should return error if poll is closed", func() {
		poll := factory.MakePoll(model.Poll{
			Status: model.POLL_CLOSED_STATUS,
		})

		err := s.pollRepo.Store(&poll)
		s.NoError(err)

		opt := factory.MakePollOption(model.PollOption{
			Content: "option 1",
			PollID:  poll.ID,
		})

		err = s.optionRepo.Store(&opt)
		s.NoError(err)

		_, err = s.svc.VoteOnPoll(VoteOnPollParams{
			PollID:       poll.ID,
			PollOptionID: opt.ID,
			UserID:       "valid-user-id",
		})

		s.Error(err)
		s.Equal(custom_err.NewInvalidActionErr("poll is not open").Error(), err.Error())
		s.Equal(len(s.publisher.Messages), 0)
	})
}
