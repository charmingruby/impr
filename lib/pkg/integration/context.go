package integration

import (
	"context"
	"time"
)

const timeout = 10

func NewContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), timeout*time.Second)
}
