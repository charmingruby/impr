package messaging

import "context"

type Subscriber interface {
	Subscribe(ctx context.Context, handler func(Message) error) error
	Close()
}
