package memory

import (
	"fmt"

	"github.com/charmingruby/impr/service/poll/internal/poll/core/model"
)

type PollOptionRepository struct {
	Items     []model.PollOption
	IsHealthy bool
}

func NewPollOptionRepository() *PollOptionRepository {
	return &PollOptionRepository{
		Items:     []model.PollOption{},
		IsHealthy: true,
	}
}

func (r *PollOptionRepository) FindByContentAndPollID(content, pollID string) (*model.PollOption, error) {
	for _, item := range r.Items {
		if item.Content == content && item.PollID == pollID {
			println("acheu")
			return &item, nil
		}
	}

	if !r.IsHealthy {
		return nil, fmt.Errorf("poll option datasource is unhealthy")
	}

	return nil, nil
}

func (r *PollOptionRepository) Store(model *model.PollOption) error {
	if !r.IsHealthy {
		return fmt.Errorf("poll option datasource is unhealthy")
	}

	r.Items = append(r.Items, *model)

	return nil
}
