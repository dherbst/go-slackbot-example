package chatty

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/websocket"
	"io/ioutil"
	"log"
	"net/http"
)

// EventHandler is an interface stored in the evHandlers map
type EventHandler interface {
	HandleEvent(c *Connection, ev *Event) error
}

type Connection struct {
	token      string
	ws         *websocket.Conn
	users      map[string]*User
	channels   map[string]*Channel
	channelIds map[string]string
	eventChan  chan *Event
	evHandlers map[string]EventHandler
	shutdown   chan bool
}

func NewConnection(token string, channel string) *Connection {
	c := &Connection{
		token:      token,
		users:      map[string]*User{},
		channels:   map[string]*Channel{},
		channelIds: map[string]string{},
		eventChan:  make(chan *Event, 100),
		evHandlers: map[string]EventHandler{},
		shutdown:   make(chan bool),
	}

	c.evHandlers["hello"] = NewHelloHandler()
	c.evHandlers["message"] = NewMessageHandler()

	return c
}

type Team struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Channel struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type RtmResponse struct {
	Ok       bool       `json:"ok"`
	Url      string     `json:"url"`
	Team     Team       `json:"team"`
	Users    []*User    `json:"users"`
	Channels []*Channel `json:"channels"`
}

type Message struct {
	User      *User    `json:"-"`
	Channel   *Channel `json:"-"`
	Text      string   `json:"text"`
	Timestamp string   `json:"ts"`
}

// Events to process
type Event struct {
	Type string `json:"Type"`
	Data string `json:"Data"` // the serialized event so you can map it later
}

// Connect establishes the websocket connection with slack.
func (c *Connection) Connect() error {
	log.Printf("Connecting to slack...\n")

	url := fmt.Sprintf("https://slack.com/api/rtm.start?token=%s", c.token)
	r, err := http.Get(url)
	if err != nil {
		log.Printf("Error rmt.start %v\n", err)
		return err
	}
	defer r.Body.Close()

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading rtm.start response %v\n", err)
		return err
	}

	var resp RtmResponse
	err = json.Unmarshal(b, &resp)
	if err != nil {
		log.Printf("Error unmarshalling rtm.start %v\n", err)
		return err
	}
	for _, user := range resp.Users {
		c.users[user.Id] = user
	}

	for _, channel := range resp.Channels {
		c.channels[channel.Id] = channel
		c.channelIds[channel.Name] = channel.Id
	}

	ws, err := websocket.Dial(resp.Url, "", "http://localhost")
	if err != nil {
		log.Printf("Error %v\n", err)
		return err
	}
	c.ws = ws

	log.Printf("Connected.\n")

	return nil
}

func (c *Connection) SendMessage(channel string, msg string) error {
	log.Printf("Sending to channel %v id %v\n", channel, c.channelIds[channel])

	payload := map[string]string{
		"type":    "message",
		"channel": c.channelIds[channel],
		"text":    msg,
	}
	return websocket.JSON.Send(c.ws, payload)
}

func (c *Connection) Close() {
	if c.ws != nil {
		c.ws.Close()
		c.ws = nil
	}
}

func (c *Connection) Run() {

	err := c.SendMessage("bot", "Started")
	if err != nil {
		log.Printf("Err from SendMessage %v\nShutting down...\n", err)
		return
	}

	// start an eventhandler
	go c.handleEvents()

	//
	c.receiveMessages()
}

// handleEvents processes a channel of events.
func (c *Connection) handleEvents() {
	for {
		ev := <-c.eventChan
		h := c.evHandlers[ev.Type]
		if h != nil {
			h.HandleEvent(c, ev)
		} else {
			log.Printf("no handler for ev=%v h=%v\n", ev, h)
		}

	}
}

// receiveMessages is the message receiving pump from the websocket
func (c *Connection) receiveMessages() {

Pump:
	for {
		select {
		case shutdown := <-c.shutdown:
			log.Printf("Got shutdown message %v\n", shutdown)
			break Pump
		default:
			slackBytes := make([]byte, 2048)
			n, err := c.ws.Read(slackBytes)
			if err != nil {
				log.Printf("Error reading %v from slack %v msg=%v",
					n,
					err,
					string(slackBytes[:n]),
				)
				continue
			}
			v := make(map[string]interface{})
			err = json.Unmarshal(slackBytes[:n], &v)
			if err != nil {
				log.Printf("Error unmarshalling %v msg=%v\n", err, slackBytes[:n])
				continue
			}
			vtype, ok := (v["type"]).(string)
			if ok {
				ev := &Event{
					Type: vtype,
					Data: string(slackBytes[:n]),
				}
				c.eventChan <- ev
			} else {
				log.Printf("Did not get string for v[type]")
			}
		}
	}
}
