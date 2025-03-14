package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewPoll(t *testing.T) {
	t.Run("it should be able to create a new Poll", func(t *testing.T) {
		in := NewPollInput{
			Title:             "Poll",
			Question:          "Is this a poll?",
			ExpirationPeriods: 1,
			OwnerID:           "owner_id",
		}

		m := NewPoll(in)

		assert.NotEmpty(t, m.ID)
		assert.Equal(t, in.Title, m.Title)
		assert.Equal(t, in.Question, m.Question)
		assert.Equal(t, POLL_OPEN_STATUS, m.Status)
		assert.Equal(t, in.ExpirationPeriods, m.ExpirationPeriods)
		assert.Equal(t, in.OwnerID, m.OwnerID)
		assert.NotZero(t, m.CreatedAt)
		assert.NotZero(t, m.UpdatedAt)
	})
}
