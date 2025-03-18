package service

import (
	"github.com/charmingruby/impr/service/poll/internal/poll/core/model"
	"github.com/charmingruby/impr/service/poll/test/factory"
)

func (s *Suite) Test_Service_GetPollSummary() {
	s.Run("should return poll summary with options and votes", func() {
		poll := factory.MakePoll(model.Poll{})
		s.summaryRepo.Polls = append(s.summaryRepo.Polls, poll)
		err := s.pollRepo.Store(&poll)
		s.NoError(err)

		opt1 := factory.MakePollOption(model.PollOption{
			Content: "option 1",
			PollID:  poll.ID,
		})
		s.summaryRepo.Options = append(s.summaryRepo.Options, opt1)

		opt2 := factory.MakePollOption(model.PollOption{
			Content: "option 2",
			PollID:  poll.ID,
		})
		s.summaryRepo.Options = append(s.summaryRepo.Options, opt2)

		vote1 := factory.MakeVote(model.Vote{
			PollID:       poll.ID,
			PollOptionID: opt1.ID,
			UserID:       "user1",
		})
		s.summaryRepo.Votes = append(s.summaryRepo.Votes, vote1)

		vote2 := factory.MakeVote(model.Vote{
			PollID:       poll.ID,
			PollOptionID: opt1.ID,
			UserID:       "user2",
		})
		s.summaryRepo.Votes = append(s.summaryRepo.Votes, vote2)

		result, err := s.svc.GetPollSummary(poll.ID)
		s.NoError(err)

		s.Equal(poll.ID, result.PollID)
		s.Equal(poll.ExpiresAt, result.ExpiresAt)
		s.Len(result.Options, 2)

		for _, option := range result.Options {
			if option.PollOptionID == opt1.ID {
				s.Equal("option 1", option.Content)
				s.Equal(2, option.Votes)
			} else if option.PollOptionID == opt2.ID {
				s.Equal("option 2", option.Content)
				s.Equal(0, option.Votes)
			}
		}
	})

	s.Run("should return an error if poll does not exist", func() {
		invalidPollID := "invalid-poll-id"

		result, err := s.svc.GetPollSummary(invalidPollID)
		s.Error(err)
		s.Nil(result)
	})
}
