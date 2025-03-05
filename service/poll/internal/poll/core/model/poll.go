package model

import (
	"time"

	"github.com/oklog/ulid/v2"
)

const (
	POLL_DRAFT_STATUS  = "draft"
	POLL_OPEN_STATUS   = "open"
	POLL_CLOSED_STATUS = "closed"
)

type Poll struct {
	ID             string    `json:"id" db:"id"`
	Name           string    `json:"name" db:"name"`
	Description    string    `json:"description" db:"description"`
	Status         string    `json:"status" db:"status"`
	ExpirationTime int       `json:"expiration_time" db:"expiration_time"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}

type NewPollInput struct {
	Name           string
	Description    string
	Status         string
	ExpirationTime int
}

func NewPoll(in NewPollInput) *Poll {
	var status string = in.Status
	if in.Status == "" {
		status = POLL_DRAFT_STATUS
	}

	return &Poll{
		ID:             ulid.Make().String(),
		Name:           in.Name,
		Description:    in.Description,
		Status:         status,
		ExpirationTime: in.ExpirationTime,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
}
