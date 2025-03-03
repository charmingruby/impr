package service

import "time"

func (s *Suite) Test_CreateAudit() {
	s.Run("should create audit", func() {
		params := CreateAuditParams{
			Context:      "context",
			Subject:      "subject",
			Content:      "content",
			DispatchedAt: time.Now(),
		}

		err := s.svc.CreateAudit(params)

		s.NoError(err)
		s.Len(s.repo.Items, 1)
		s.Equal(params.Context, s.repo.Items[0].Context)
		s.Equal(params.Subject, s.repo.Items[0].Subject)
		s.Equal(params.Content, s.repo.Items[0].Content)
		s.Equal(params.DispatchedAt, s.repo.Items[0].DispatchedAt)
	})
}
