package service

type RefreshSessionParams struct {
	RefreshToken string
}

type RefreshSessionResult struct {
	RenewedAccessToken string
}

func (s *Service) RefreshSession(in RefreshSessionParams) (RefreshSessionResult, error) {
	token, err := s.identityProvider.RefreshSession(in.RefreshToken)
	if err != nil {
		return RefreshSessionResult{}, err
	}

	return RefreshSessionResult{
		RenewedAccessToken: token,
	}, nil
}
