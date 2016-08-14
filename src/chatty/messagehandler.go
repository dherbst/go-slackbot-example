package chatty

import (
	"encoding/json"
	"errors"
	"log"
)

type MessageHandler struct {
	Type string
}

func NewMessageHandler() *MessageHandler {
	h := &MessageHandler{Type: "message"}
	return h
}

// HandleEvent handles the message event passed in.
func (h *MessageHandler) HandleEvent(c *Connection, ev *Event) error {
	if ev.Type != h.Type {
		log.Printf("%v handler got wrong type %v\n", h.Type, ev.Type)
		return errors.New("Wrong event type")
	}

	data := make(map[string]interface{})
	err := json.Unmarshal([]byte(ev.Data), &data)
	if err != nil {
		log.Printf("Erorr unmarshalling ev.Data %v\n", err)
		return err
	}

	log.Printf("Got message event %v\n", data["text"])
	return nil
}
