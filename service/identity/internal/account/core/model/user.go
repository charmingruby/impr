package model

import (
	"time"

	"github.com/oklog/ulid/v2"
)

type User struct {
	ID        string     `json:"id" db:"id"`
	Name      string     `json:"name" db:"name"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
}

type NewUserInput struct {
	Name string
}

func NewUser(in NewUserInput) *User {
	return &User{
		ID:        ulid.Make().String(),
		Name:      in.Name,
		CreatedAt: time.Now(),
		UpdatedAt: nil,
		DeletedAt: nil,
	}
}

func FromUser(in User) *User {
    return &User{
        ID:        in.ID,
        Name:      in.Name,
        CreatedAt: in.CreatedAt,
        UpdatedAt: in.UpdatedAt,
        DeletedAt: in.DeletedAt,
    }
}

func (m *User) SoftDelete() {
	now := time.Now()

	m.UpdatedAt = &now
	m.DeletedAt = &now
}
