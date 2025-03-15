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
	ID                string    `json:"id" db:"id"`
	Title             string    `json:"title" db:"title"`
	Question          string    `json:"question" db:"question"`
	Status            string    `json:"status" db:"status"`
	ExpirationPeriods int       `json:"expiration_periods" db:"expiration_periods"`
	OwnerID           string    `json:"owner_id" db:"owner_id"`
	CreatedAt         time.Time `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time `json:"updated_at" db:"updated_at"`
}

type NewPollInput struct {
	Title             string
	Question          string
	ExpirationPeriods int
	OwnerID           string
}

func NewPoll(in NewPollInput) *Poll {
	return &Poll{
		ID:                ulid.Make().String(),
		Title:             in.Title,
		Question:          in.Question,
		Status:            POLL_OPEN_STATUS,
		ExpirationPeriods: in.ExpirationPeriods,
		OwnerID:           in.OwnerID,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}
}
