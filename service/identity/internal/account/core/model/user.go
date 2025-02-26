package model

import (
	"encoding/json"
	"time"
)

type User struct {
	ID         string     `json:"id"`
	FirstName  string     `json:"first_name"`
	LastName   string     `json:"last_name"`
	Email      string     `json:"email"`
	IsVerified bool       `json:"is_verified"`
	Birthdate  time.Time  `json:"birthdate"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
}

func (u *User) Marshal() ([]byte, error) {
	return json.Marshal(u)
}

func UnmarshalUser(payload []byte) (*User, error) {
	u := User{}

	if err := json.Unmarshal(payload, &u); err != nil {
		return nil, err
	}

	return &u, nil
}
