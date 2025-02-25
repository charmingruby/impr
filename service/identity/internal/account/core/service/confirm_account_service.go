package service

import "github.com/charmingruby/impr/service/identity/internal/account/core/gateway"

type ConfirmAccountParams struct {
	Email string
	Code  string
}

func (s *Service) ConfirmAccount(in ConfirmAccountParams) error {
	return s.identityProvider.ConfirmAccount(gateway.ConfirmAccountInput{
		Email: in.Email,
		Code:  in.Code,
	})
}
