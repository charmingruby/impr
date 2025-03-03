package model

import (
	"encoding/json"
	"time"
)

type Audit struct {
	ID           string    `json:"id" bson:"_id"`
	Context      string    `json:"context" bson:"context"`
	Subject      string    `json:"subject" bson:"subject"`
	Content      string    `json:"content" bson:"content"`
	DispatchedAt time.Time `json:"dispatched_at" bson:"dispatched_at"`
	CreatedAt    time.Time `json:"created_at" bson:"created_at"`
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
