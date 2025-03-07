package service

import (
	"testing"

	"github.com/charmingruby/impr/service/poll/internal/poll/core/model"
	"github.com/charmingruby/impr/service/poll/test/memory"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite

	pollRepo   *memory.PollRepository
	optionRepo *memory.PollOptionRepository
	svc        *Service
}

func (s *Suite) SetupTest() {
	s.pollRepo = memory.NewPollRepository()
	s.optionRepo = memory.NewPollOptionRepository()
	s.svc = New(s.pollRepo, s.optionRepo)
}

func (s *Suite) SetupSubTest() {
	s.pollRepo.Items = []model.Poll{}
	s.pollRepo.IsHealthy = true

	s.optionRepo.Items = []model.PollOption{}
	s.optionRepo.IsHealthy = true
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}
