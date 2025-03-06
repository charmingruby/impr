package model

import (
	"time"

	"github.com/oklog/ulid/v2"
)

const (
	POLL_OPEN_STATUS   = "open"
	POLL_CLOSED_STATUS = "closed"
)

type Poll struct {
	ID             string    `json:"id" db:"id"`
	Name           string    `json:"name" db:"name"`
	Description    string    `json:"description" db:"description"`
	Status         string    `json:"status" db:"status"`
	ExpirationTime int       `json:"expiration_time" db:"expiration_time"`
	OwnerID        string    `json:"owner_id" db:"owner_id"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}

type NewPollInput struct {
	Name           string
	Description    string
	ExpirationTime int
	OwnerID        string
}

func NewPoll(in NewPollInput) *Poll {
	return &Poll{
		ID:             ulid.Make().String(),
		Name:           in.Name,
		Description:    in.Description,
		Status:         POLL_OPEN_STATUS,
		ExpirationTime: in.ExpirationTime,
		OwnerID:        in.OwnerID,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
}
