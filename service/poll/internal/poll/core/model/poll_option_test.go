package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewPollOption(t *testing.T) {
	t.Run("it should be able to create a new PollOption", func(t *testing.T) {
		in := NewPollOptionInput{
			Content: "Poll Option",
			PollID:  "poll-id",
		}

		m := NewPollOption(in)

		assert.NotEmpty(t, m.ID)
		assert.Equal(t, in.Content, m.Content)
		assert.Equal(t, in.PollID, m.PollID)
		assert.NotZero(t, m.CreatedAt)
	})
}
