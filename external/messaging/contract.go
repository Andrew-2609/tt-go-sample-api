package messaging

import "context"

// Message represents a message to be produced
// to an external or internal message producer.
type Message struct {
	// ID is the unique identifier of the message.
	ID string

	// GroupID is an information that can be shared
	// by multiple messages of the same category (e.g.
	// messages regarding a same client by its document.)
	GroupID string

	// Payload is the actual Message payload, that's going
	// to be serialized and sent.
	Payload any
}

// MessageProducer is an interface that any
// message producer within the application must
// implement in order to produce messages to
// an external or internal service.
type MessageProducer interface {
	Produce(ctx context.Context, message Message) error
}
