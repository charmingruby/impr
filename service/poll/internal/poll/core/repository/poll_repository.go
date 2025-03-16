package repository

import (
	"time"

	"github.com/charmingruby/impr/service/poll/internal/poll/core/model"
)

type PollRepository interface {
	FindByID(id string) (*model.Poll, error)
	FindByIDAndOwnerID(id, ownerID string) (*model.Poll, error)
	FindByTitleAndOwnerID(title, ownerID string) (*model.Poll, error)
	FindAllWithInferiorExpiresAt(expiresAt time.Time) ([]model.Poll, error)
	Store(model *model.Poll) error
	Save(model *model.Poll) error
}
