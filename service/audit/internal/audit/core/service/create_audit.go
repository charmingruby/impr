package service

import (
	"time"

	"github.com/charmingruby/impr/service/audit/internal/audit/core/model"
	"github.com/charmingruby/impr/service/audit/internal/shared/id"
)

type CreateAuditParams struct {
	Context      string    `json:"context"`
	Subject      string    `json:"subject"`
	Content      string    `json:"content"`
	DispatchedAt time.Time `json:"dispatched_at"`
}

func (s *Service) CreateAudit(in CreateAuditParams) error {
	audit := model.Audit{
		ID:           id.New(),
		Context:      in.Context,
		Subject:      in.Subject,
		Content:      in.Content,
		DispatchedAt: in.DispatchedAt,
		CreatedAt:    time.Now(),
	}

	return s.repo.Create(audit)
}
