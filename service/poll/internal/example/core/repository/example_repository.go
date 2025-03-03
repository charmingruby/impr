package repository

import "github.com/charmingruby/bob/internal/example/core/model"

type ExampleRepository interface {
	Store(model *model.Example) error
	FindByID(id string) (*model.Example, error)
	Delete(model *model.Example) error
}
