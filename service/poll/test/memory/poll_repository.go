package memory

import (
	"fmt"

	"github.com/charmingruby/impr/service/poll/internal/poll/core/model"
)

type PollRepository struct {
	Items     []model.Poll
	IsHealthy bool
}

func NewPollRepository() *PollRepository {
	return &PollRepository{
		Items:     []model.Poll{},
		IsHealthy: true,
	}
}

func (r *PollRepository) FindByID(id string) (*model.Poll, error) {
	for _, item := range r.Items {
		if item.ID == id {
			return &item, nil
		}
	}

	if !r.IsHealthy {
		return nil, fmt.Errorf("poll datasource is unhealthy")
	}

	return nil, nil
}

func (r *PollRepository) FindByTitleAndOwnerID(title, ownerID string) (*model.Poll, error) {
	for _, item := range r.Items {
		if item.Title == title && item.OwnerID == ownerID {
			return &item, nil
		}
	}

	if !r.IsHealthy {
		return nil, fmt.Errorf("poll datasource is unhealthy")
	}

	return nil, nil
}

func (r *PollRepository) FindByIDAndOwnerID(id, ownerID string) (*model.Poll, error) {
	for _, item := range r.Items {
		if item.ID == id && item.OwnerID == ownerID {
			return &item, nil
		}
	}

	if !r.IsHealthy {
		return nil, fmt.Errorf("poll datasource is unhealthy")
	}

	return nil, nil
}

func (r *PollRepository) Store(model *model.Poll) error {
	if !r.IsHealthy {
		return fmt.Errorf("poll datasource is unhealthy")
	}

	r.Items = append(r.Items, *model)

	return nil
}

func (r *PollRepository) Save(model *model.Poll) error {
	for i, item := range r.Items {
		if item.ID == model.ID {
			r.Items[i] = *model
			return nil
		}
	}

	if !r.IsHealthy {
		return fmt.Errorf("poll datasource is unhealthy")
	}

	return nil
}
