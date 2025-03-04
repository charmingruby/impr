package model

import (
	"time"

	"github.com/oklog/ulid/v2"
)

type Poll struct {
	ID        string     `json:"id" db:"id"`
	Name      string     `json:"name" db:"name"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
}

type NewPollInput struct {
	Name string
}

func NewPoll(in NewPollInput) *Poll {
	return &Poll{
		ID:        ulid.Make().String(),
		Name:      in.Name,
		CreatedAt: time.Now(),
		UpdatedAt: nil,
		DeletedAt: nil,
	}
}

func FromPoll(in Poll) *Poll {
    return &Poll{
        ID:        in.ID,
        Name:      in.Name,
        CreatedAt: in.CreatedAt,
        UpdatedAt: in.UpdatedAt,
        DeletedAt: in.DeletedAt,
    }
}

func (m *Poll) SoftDelete() {
	now := time.Now()

	m.UpdatedAt = &now
	m.DeletedAt = &now
}
