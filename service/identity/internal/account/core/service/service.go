package service

import "github.com/charmingruby/impr/service/identity/internal/account/core/repository"

type Service struct {
	userRepository repository.UserRepository
}

type Input struct {
	UserRepository repository.UserRepository
}

func New(in Input) *Service {
	return &Service{
		userRepository: in.UserRepository,
	}
}
