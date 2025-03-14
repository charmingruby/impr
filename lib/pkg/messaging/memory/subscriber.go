package memory

import (
	"context"
	"errors"

	"github.com/charmingruby/impr/lib/pkg/messaging"
)

type Subscriber struct {
	Messages  []messaging.Message
	IsHealthy bool
}

func NewSubscriber() *Subscriber {
	return &Subscriber{
		Messages:  []messaging.Message{},
		IsHealthy: true,
	}
}

func (p *Subscriber) Subscribe(ctx context.Context, handler func(messaging.Message) error) error {
	if !p.IsHealthy {
		return errors.New("subscriber is not healthy")
	}

	for _, message := range p.Messages {
		if err := handler(message); err != nil {
			return err
		}
	}

	return nil
}

func (p *Subscriber) Close() {
	p.IsHealthy = false
	p.Messages = []messaging.Message{}
}
