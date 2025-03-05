package model

import (
	"time"

	"github.com/oklog/ulid/v2"
)

type Vote struct {
	ID           string    `json:"id" db:"id"`
	PollOptionID string    `json:"poll_option_id" db:"poll_option_id"`
	UserID       string    `json:"user_id" db:"user_id"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}

type NewVoteInput struct {
	PollOptionID string
	UserID       string
}

func NewVote(in NewVoteInput) *Vote {
	return &Vote{
		ID:           ulid.Make().String(),
		PollOptionID: in.PollOptionID,
		UserID:       in.UserID,
		CreatedAt:    time.Now(),
	}
}
