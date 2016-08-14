package chatty

import (
	"errors"
	"log"
)

type HelloHandler struct {
	Type string
}

func NewHelloHandler() *HelloHandler {
	h := &HelloHandler{
		Type: "hello",
	}
	return h
}

// HandleEvent handles the Event passed in
func (h *HelloHandler) HandleEvent(c *Connection, ev *Event) error {
	if ev.Type != h.Type {
		log.Printf("HelloHandler got wrong type %v\n", ev.Type)
		return errors.New("WrongType")
	}

	log.Printf("Got hello event\n")
	return nil
}
