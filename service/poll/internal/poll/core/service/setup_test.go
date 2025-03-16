package service

import (
	"testing"

	"github.com/charmingruby/impr/lib/pkg/messaging"
	msgMemory "github.com/charmingruby/impr/lib/pkg/messaging/memory"
	"github.com/charmingruby/impr/service/poll/internal/poll/core/model"
	"github.com/charmingruby/impr/service/poll/pkg/logger"
	"github.com/charmingruby/impr/service/poll/test/memory"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite

	pollRepo   *memory.PollRepository
	optionRepo *memory.PollOptionRepository
	voteRepo   *memory.VoteRepository
	publisher  *msgMemory.Publisher
	svc        *Service
}

func (s *Suite) SetupTest() {
	logger.New()
	s.pollRepo = memory.NewPollRepository()
	s.optionRepo = memory.NewPollOptionRepository()
	s.voteRepo = memory.NewVoteRepository()
	s.publisher = msgMemory.NewPublisher()
	s.svc = New(s.pollRepo, s.optionRepo, s.voteRepo, s.publisher)
}

func (s *Suite) SetupSubTest() {
	s.pollRepo.Items = []model.Poll{}
	s.pollRepo.IsHealthy = true

	s.optionRepo.Items = []model.PollOption{}
	s.optionRepo.IsHealthy = true

	s.voteRepo.Items = []model.Vote{}
	s.voteRepo.IsHealthy = true

	s.publisher.IsHealthy = true
	s.publisher.Messages = []messaging.Message{}
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}
