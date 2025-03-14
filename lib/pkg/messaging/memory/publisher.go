package memory

import (
	"errors"

	"github.com/charmingruby/impr/lib/pkg/messaging"
)

type Publisher struct {
	Messages  []messaging.Message
	IsHealthy bool
}

func NewPublisher() *Publisher {
	return &Publisher{
		Messages:  []messaging.Message{},
		IsHealthy: true,
	}
}

func (p *Publisher) Publish(message messaging.Message) error {
	if p.IsHealthy {
		p.Messages = append(p.Messages, message)

		return nil
	}

	return errors.New("publisher is not healthy")
}

func (p *Publisher) Close() {
	p.IsHealthy = false
	p.Messages = []messaging.Message{}
}
