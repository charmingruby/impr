package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewVote(t *testing.T) {
	t.Run("it should be able to create a new Vote", func(t *testing.T) {
		in := NewVoteInput{
			PollOptionID: "option-id",
			UserID:       "user-id",
		}

		m := NewVote(in)

		assert.NotEmpty(t, m.ID)
		assert.Equal(t, in.PollOptionID, m.PollOptionID)
		assert.Equal(t, in.UserID, m.UserID)
		assert.NotZero(t, m.CreatedAt)
	})
}
