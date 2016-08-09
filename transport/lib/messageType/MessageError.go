package messageType

import (
	"log"
	"github.com/iain17/goTransport/transport/lib/Message"
)

type messageError struct {
	Message.Message
	Reason error `json:"reason"`
}

func init() {
	Message.Set(NewMessageError(nil))
}

func NewMessageError(reason error) *messageError {
	return &messageError{
		Message: Message.NewMessage(Message.MessageTypeError),
		Reason: reason,
	}
}

func (message messageError) Sending() error {
	return nil
}

func (message messageError) Received() error {
	log.Print(message.Reason)
	return 	nil
}