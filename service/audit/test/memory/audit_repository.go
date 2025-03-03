package memory

import (
	"github.com/charmingruby/impr/service/audit/internal/audit/core/model"
)

type AuditRepository struct {
	Items []model.Audit
}

func NewAuditRepository() *AuditRepository {
	return &AuditRepository{
		Items: []model.Audit{},
	}
}

func (r *AuditRepository) Create(audit model.Audit) error {
	r.Items = append(r.Items, audit)

	return nil
}
