package service

import (
	"github.com/charmingruby/impr/service/identity/internal/account/core/gateway"
)

type SignInParams struct {
	FirstName string
	LastName  string
	Email     string
	Birthdate string
	Password  string
}

type SignInResult struct {
	RefreshToken string
	AccessToken  string
}

func (s *Service) SignIn(in SignInParams) (SignInResult, error) {
	result, err := s.identityProvider.SignIn(gateway.SignInInput{
		Email:    in.Email,
		Password: in.Password,
	})
	if err != nil {
		return SignInResult{}, err
	}

	return SignInResult{
		RefreshToken: result.RefreshToken,
		AccessToken:  result.AccessToken,
	}, nil
}
