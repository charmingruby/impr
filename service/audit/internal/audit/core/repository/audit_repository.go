package repository

import "github.com/charmingruby/impr/service/audit/internal/audit/core/model"

type Filter struct {
	ID      string
	Context string
	Subject string
}

type AuditRepository interface {
	Create(audit model.Audit) error
	//FindByFilter(filter Filter) ([]model.Audit, error)
}
