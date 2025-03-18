package model

import (
	"time"
)

type PollSummaryOption struct {
	PollOptionID string
	Content      string
	Votes        int
}

type PollSummary struct {
	PollID    string
	Options   []PollSummaryOption
	ExpiresAt *time.Time
}
