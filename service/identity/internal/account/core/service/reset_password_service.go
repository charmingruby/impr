package service

import "github.com/charmingruby/impr/service/identity/internal/account/core/gateway"

type ResetPasswordParams struct {
	Email            string
	ConfirmationCode string
	NewPassword      string
}

func (s *Service) ResetPassword(in ResetPasswordParams) error {
	return s.identityProvider.ResetPassword(gateway.ResetPasswordInput{
		Email:            in.Email,
		ConfirmationCode: in.ConfirmationCode,
		NewPassword:      in.NewPassword,
	})
}
