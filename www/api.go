package slackbot

import (
	"fmt"
	"net/http"
)

type SlackCommand struct {
	ChannelId   string
	ChannelName string
	UserId      string
	UserName    string
	Command     string
	TeamId      string
	TeamDomain  string
	Text        string
	Token       string
}

// UnMarshalCommand takes the request from slack, returns a SlackCommand object
func UnMarshalCommand(r *http.Request) (*SlackCommand, error) {
	c := &SlackCommand{}

	c.ChannelId = r.FormValue("channel_id")
	c.ChannelName = r.FormValue("channel_name")
	c.UserId = r.FormValue("user_id")
	c.UserName = r.FormValue("user_name")
	c.Command = r.FormValue("command")
	c.TeamId = r.FormValue("team_id")
	c.TeamDomain = r.FormValue("team_domain")
	c.Text = r.FormValue("text")
	c.Token = r.FormValue("token")

	return c, nil
}

// Dispatch the command based on the parameter
func CommandHandler(w http.ResponseWriter, r *http.Request) {

	command, err := UnMarshalCommand(r)
	if err != nil {
		http.Error(w, "Could not process command", http.StatusBadRequest)
	}

	fmt.Fprintf(w, "Hello there command=%v text=%v", command.Command, command.Text)
}
