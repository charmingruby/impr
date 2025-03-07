package repository

import "github.com/charmingruby/impr/service/poll/internal/poll/core/model"

type PollRepository interface {
	FindByID(id string) (*model.Poll, error)
	FindByTitleAndOwnerID(title, ownerID string) (*model.Poll, error)
	Store(model *model.Poll) error
	Save(model *model.Poll) error
}
