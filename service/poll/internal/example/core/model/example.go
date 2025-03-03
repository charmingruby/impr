package model

import (
	"time"

	"github.com/oklog/ulid/v2"
)

type Example struct {
	ID        string     `json:"id" db:"id"`
	Name      string     `json:"name" db:"name"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
}

type NewExampleInput struct {
	Name string
}

func NewExample(in NewExampleInput) *Example {
	return &Example{
		ID:        ulid.Make().String(),
		Name:      in.Name,
		CreatedAt: time.Now(),
		UpdatedAt: nil,
		DeletedAt: nil,
	}
}

func FromExample(in Example) *Example {
    return &Example{
        ID:        in.ID,
        Name:      in.Name,
        CreatedAt: in.CreatedAt,
        UpdatedAt: in.UpdatedAt,
        DeletedAt: in.DeletedAt,
    }
}

func (m *Example) SoftDelete() {
	now := time.Now()

	m.UpdatedAt = &now
	m.DeletedAt = &now
}
