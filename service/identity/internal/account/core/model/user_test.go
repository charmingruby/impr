package model

import (
    "testing"
    "time"

    "github.com/stretchr/testify/assert"
    "github.com/oklog/ulid/v2"
)

func Test_NewUser(t *testing.T) {
    t.Run("it should be able to create a new User", func(t *testing.T) {
        in := NewUserInput{
            Name: "User",
        }
        m := NewUser(in)

        assert.NotEmpty(t, m.ID)
        assert.Equal(t, in.Name, m.Name)
        assert.NotZero(t, m.CreatedAt)
        assert.Nil(t, m.UpdatedAt)
        assert.Nil(t, m.DeletedAt)
    })
}

func Test_FromUser(t *testing.T) {
    t.Run("it should be able to create a User from input", func(t *testing.T) {
        in := User{
            ID:        ulid.Make().String(),
            Name:      "User",
            CreatedAt: time.Now(),
            UpdatedAt: nil,
            DeletedAt: nil,
        }
        m := FromUser(in)

        assert.Equal(t, in.ID, m.ID)
        assert.Equal(t, in.Name, m.Name)
        assert.Equal(t, in.CreatedAt, m.CreatedAt)
        assert.Equal(t, in.UpdatedAt, m.UpdatedAt)
        assert.Equal(t, in.DeletedAt, m.DeletedAt)
    })
}

func Test_SoftDeleteUser(t *testing.T) {
    t.Run("it should be able to soft delete a User", func(t *testing.T) {
        now := time.Now()

        in := User{
            ID:        ulid.Make().String(),
            Name:      "User",
            CreatedAt: now,
            UpdatedAt: nil,
            DeletedAt: nil,
        }
        
        m := FromUser(in)

        m.SoftDelete()

        assert.NotNil(t, m.DeletedAt)
        assert.NotNil(t, m.UpdatedAt)
        assert.True(t, m.DeletedAt.After(now) || m.DeletedAt.Equal(now))
        assert.True(t, m.UpdatedAt.After(now) || m.UpdatedAt.Equal(now))
    })
}