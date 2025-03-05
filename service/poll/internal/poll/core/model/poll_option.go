package model

import (
	"time"

	"github.com/oklog/ulid/v2"
)

type PollOption struct {
	ID        string    `json:"id" db:"id"`
	Content   string    `json:"content" db:"content"`
	PollID    string    `json:"poll_id" db:"poll_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type NewPollOptionInput struct {
	Content string
	PollID  string
}

func NewPollOption(in NewPollOptionInput) *PollOption {
	return &PollOption{
		ID:        ulid.Make().String(),
		Content:   in.Content,
		PollID:    in.PollID,
		CreatedAt: time.Now(),
	}
}
