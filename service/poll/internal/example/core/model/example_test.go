package model

import (
    "testing"
    "time"

    "github.com/stretchr/testify/assert"
    "github.com/oklog/ulid/v2"
)

func Test_NewExample(t *testing.T) {
    t.Run("it should be able to create a new Example", func(t *testing.T) {
        in := NewExampleInput{
            Name: "Example",
        }
        m := NewExample(in)

        assert.NotEmpty(t, m.ID)
        assert.Equal(t, in.Name, m.Name)
        assert.NotZero(t, m.CreatedAt)
        assert.Nil(t, m.UpdatedAt)
        assert.Nil(t, m.DeletedAt)
    })
}

func Test_FromExample(t *testing.T) {
    t.Run("it should be able to create a Example from input", func(t *testing.T) {
        in := Example{
            ID:        ulid.Make().String(),
            Name:      "Example",
            CreatedAt: time.Now(),
            UpdatedAt: nil,
            DeletedAt: nil,
        }
        m := FromExample(in)

        assert.Equal(t, in.ID, m.ID)
        assert.Equal(t, in.Name, m.Name)
        assert.Equal(t, in.CreatedAt, m.CreatedAt)
        assert.Equal(t, in.UpdatedAt, m.UpdatedAt)
        assert.Equal(t, in.DeletedAt, m.DeletedAt)
    })
}

func Test_SoftDeleteExample(t *testing.T) {
    t.Run("it should be able to soft delete a Example", func(t *testing.T) {
        now := time.Now()

        in := Example{
            ID:        ulid.Make().String(),
            Name:      "Example",
            CreatedAt: now,
            UpdatedAt: nil,
            DeletedAt: nil,
        }
        
        m := FromExample(in)

        m.SoftDelete()

        assert.NotNil(t, m.DeletedAt)
        assert.NotNil(t, m.UpdatedAt)
        assert.True(t, m.DeletedAt.After(now) || m.DeletedAt.Equal(now))
        assert.True(t, m.UpdatedAt.After(now) || m.UpdatedAt.Equal(now))
    })
}