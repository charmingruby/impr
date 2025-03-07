package service

import (
	"testing"

	"github.com/charmingruby/impr/service/audit/internal/audit/core/model"
	"github.com/charmingruby/impr/service/audit/test/memory"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite

	repo memory.AuditRepository
	svc  *Service
}

func (s *Suite) SetupTest() {
	s.repo = *memory.NewAuditRepository()
	s.svc = New(&s.repo)
}

func (s *Suite) SetupSubTest() {
	s.repo.Items = []model.Audit{}
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}
