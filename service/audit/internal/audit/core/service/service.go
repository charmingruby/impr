package service

import "github.com/charmingruby/impr/service/audit/internal/audit/core/repository"

type Service struct {
	repo repository.AuditRepository
}

func New(repo repository.AuditRepository) *Service {
	return &Service{
		repo: repo,
	}
}
