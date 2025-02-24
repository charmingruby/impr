package service

import "github.com/charmingruby/impr/service/identity/internal/account/core/gateway"

type Service struct {
	identityProvider gateway.IdentityProvider
}

func New(identityProvider gateway.IdentityProvider) *Service {
	return &Service{
		identityProvider: identityProvider,
	}
}
