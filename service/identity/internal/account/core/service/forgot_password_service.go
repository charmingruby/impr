package service

type ForgotPasswordParams struct {
	Email string
}

func (s *Service) ForgotPassword(in ForgotPasswordParams) error {
	return s.identityProvider.ForgotPassword(in.Email)
}
