package repository

import "github.com/charmingruby/impr/service/identity/internal/account/core/model"

type UserRepository interface {
	Store(model *model.User) error
	FindByID(id string) (*model.User, error)
	Delete(model *model.User) error
}
