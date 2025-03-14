package messaging

type Publisher interface {
	Publish(message Message) error
	Close()
}
