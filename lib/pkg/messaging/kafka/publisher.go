package kafka

import (
	"fmt"

	"github.com/charmingruby/impr/lib/pkg/messaging"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type Publisher struct {
	producer *kafka.Producer
	topic    string
}

func NewPublisher(brokerURL, topic string) (*Publisher, error) {
	config := &kafka.ConfigMap{
		"bootstrap.servers": brokerURL,
	}

	producer, err := kafka.NewProducer(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create Kafka producer: %w", err)
	}

	return &Publisher{
		producer: producer,
		topic:    topic,
	}, nil
}

func (p *Publisher) Publish(message messaging.Message) error {
	deliveryChan := make(chan kafka.Event, 1)
	defer close(deliveryChan)

	err := p.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &p.topic, Partition: kafka.PartitionAny},
		Key:            []byte(message.Key),
		Value:          message.Value,
	}, deliveryChan)

	if err != nil {
		return fmt.Errorf("failed to produce message: %w", err)
	}

	e := <-deliveryChan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		return fmt.Errorf("failed to deliver message: %w", m.TopicPartition.Error)
	}

	return nil
}

func (p *Publisher) Close() {
	p.producer.Flush(15 * 1000)
	p.producer.Close()
}
