package Message

import (
	"log"
	"reflect"
	"encoding/json"
)

func build(definition reflect.Type, data string) IMessage {
	_message := reflect.New(definition).Interface()
	err := json.Unmarshal([]byte(data), &_message)
	if err != nil {
		log.Print("Error unmarshalling the message", err)
		return nil
	}
	if message, ok := _message.(IMessage); ok {
		return message
	}
	log.Print("Could not cast message to IMessage interface. Invalid MessageType")
	return nil
}