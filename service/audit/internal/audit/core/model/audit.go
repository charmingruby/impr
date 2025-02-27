package model

import (
	"encoding/json"
	"time"
)

type Audit struct {
	ID           string    `json:"id"`
	Context      string    `json:"context"`
	Subject      string    `json:"subject"`
	Content      string    `json:"content"`
	DispatchedAt time.Time `json:"dispatched_at"`
	CreatedAt    time.Time `json:"created_at"`
}

func (a *Audit) Marshal() ([]byte, error) {
	return json.Marshal(a)
}

func UnmarshalAudit(payload []byte) (*Audit, error) {
	a := Audit{}

	if err := json.Unmarshal(payload, &a); err != nil {
		return nil, err
	}

	return &a, nil
}
