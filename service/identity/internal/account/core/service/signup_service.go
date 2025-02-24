package service

import (
	"github.com/charmingruby/impr/lib/pkg/errs"
	"github.com/charmingruby/impr/service/identity/internal/account/core/gateway"
	"github.com/charmingruby/impr/service/identity/pkg/helper"
)

type SignUpParams struct {
	FirstName string
	LastName  string
	Email     string
	Birthdate string
	Password  string
}

func (s *Service) SignUp(in SignUpParams) (string, error) {
	parsedBirthdate, err := helper.StringToBirthdate(in.Birthdate)
	if err != nil {
		return "", errs.NewInvalidFieldFormatErr("birthdate", err)
	}

	userID, err := s.identityProvider.SignUp(gateway.SignUpInput{
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Email:     in.Email,
		Birthdate: parsedBirthdate,
		Password:  in.Password,
	})
	if err != nil {
		return "", err
	}

	return userID, nil
}
