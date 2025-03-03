package kafka

import (
	"context"
	"fmt"
	"log"

	"github.com/charmingruby/impr/lib/pkg/messaging"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type Subscriber struct {
	subscriber *kafka.Consumer
}

func NewSubscriber(brokerURL, topic, groupID string) (*Subscriber, error) {
	config := &kafka.ConfigMap{
		"bootstrap.servers":  brokerURL,
		"group.id":           groupID,
		"auto.offset.reset":  "earliest",
		"enable.auto.commit": false,
	}

	subscriber, err := kafka.NewConsumer(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create Kafka consumer: %w", err)
	}

	err = subscriber.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to subscribe to topic: %w", err)
	}

	return &Subscriber{
		subscriber: subscriber,
	}, nil
}

func (k *Subscriber) Subscribe(ctx context.Context, handler func(messaging.Message) error) error {
	for {
		select {
		case <-ctx.Done():
			log.Println("context cancelled, stopping consumer")
			return nil

		default:
			msg, err := k.subscriber.ReadMessage(-1)
			if err != nil {
				if err.(kafka.Error).Code() == kafka.ErrTimedOut {
					continue
				}
				return fmt.Errorf("failed to read message: %w", err)
			}

			message := messaging.Message{
				Key:   fmt.Sprintf("%v", msg.Key),
				Value: msg.Value,
			}

			if err := handler(message); err != nil {
				return fmt.Errorf("handler failed: %w", err)
			}

			_, err = k.subscriber.CommitMessage(msg)
			if err != nil {
				return fmt.Errorf("failed to commit message: %w", err)
			}
		}
	}
}

func (k *Subscriber) Close() {
	if k.subscriber != nil {
		k.subscriber.Close()
	}
}
