package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_User_Marshalling(t *testing.T) {
	t.Run("it should be able to marshal and unmarshal user model", func(t *testing.T) {
		birthdate, err := time.Parse("01-02-2006", "05-31-2003")
		assert.NoError(t, err)

		u := User{
			ID:         "id",
			FirstName:  "john",
			LastName:   "doe",
			Email:      "john@doe.com",
			IsVerified: false,
			Birthdate:  birthdate,
			CreatedAt:  time.Now(),
			UpdatedAt:  nil,
		}

		payload, err := u.Marshal()
		assert.NoError(t, err)

		unmarshaledUser, err := UnmarshalUser(payload)

		assert.NoError(t, err)
		assert.Equal(t, u, *unmarshaledUser)
	})
}
