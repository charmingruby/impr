package repository

import "github.com/charmingruby/impr/service/audit/internal/audit/core/model"

type AuditRepository interface {
	Create(audit model.Audit) error
}
