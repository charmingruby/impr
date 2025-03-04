package repository

import "github.com/charmingruby/impr/service/poll/internal/poll/core/model"

type PollRepository interface {
	Store(model *model.Poll) error
	FindByID(id string) (*model.Poll, error)
	Delete(model *model.Poll) error
}
