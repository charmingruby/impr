package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_Audit_Marshalling(t *testing.T) {
	t.Run("it should be able to marshal and unmarshal audit model", func(t *testing.T) {
		a := Audit{
			ID:           "id",
			Context:      "context",
			Subject:      "subject",
			Content:      "content",
			DispatchedAt: time.Now(),
			CreatedAt:    time.Now(),
		}

		payload, err := a.Marshal()
		assert.NoError(t, err)

		unmarshaledAudit, err := UnmarshalAudit(payload)

		assert.NoError(t, err)

		assert.Equal(t, a.ID, unmarshaledAudit.ID)
		assert.Equal(t, a.Context, unmarshaledAudit.Context)
		assert.Equal(t, a.Subject, unmarshaledAudit.Subject)
		assert.Equal(t, a.Content, unmarshaledAudit.Content)
		assert.WithinDuration(t, a.DispatchedAt, unmarshaledAudit.DispatchedAt, time.Second)
		assert.WithinDuration(t, a.CreatedAt, unmarshaledAudit.CreatedAt, time.Second)
	})
}
