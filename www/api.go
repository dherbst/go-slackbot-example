package slackbot

import (
	"fmt"
	"net/http"
)

// SlackCommand is a struct holding the values that slack will post to our bot
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

// SlackResult holds the result of processing the command.
type SlackResult struct {
	IsTextResult bool
	Text         string
	Command      *SlackCommand
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
		return
	}

	result, err := ProcessCommand(command)
	if err != nil {
		http.Error(w, "Error processing result", http.StatusInternalServerError)
		return
	}
	if result.IsTextResult {
		fmt.Fprintf(w, "%v", result.Text)
	} else {
		err = SendResult(result)
		if err != nil {
			http.Error(w, "Error processing result", http.StatusInternalServerError)
			return
		}
		http.Error(w, "", http.StatusNoContent) // return a 204
	}
}

// Determine what do to with the different text commands
func ProcessCommand(cmd *SlackCommand) (*SlackResult, error) {
	var result *SlackResult
	var err error
	switch cmd.Text {

	case "hello":
		result, err = HelloCommand(cmd)

	default:
		result, err = UnknownCommand(cmd)
	}
	return result, err
}

func HelloCommand(cmd *SlackCommand) (*SlackResult, error) {
	result := &SlackResult{
		IsTextResult: true,
		Text:         fmt.Sprintf("hi %v", cmd.UserName),
		Command:      cmd,
	}

	return result, nil
}

func UnknownCommand(cmd *SlackCommand) (*SlackResult, error) {
	result := &SlackResult{
		IsTextResult: true,
		Text:         "I don't know that command, sorry",
		Command:      cmd,
	}
	return result, nil
}

// SendResult sends the result.Text to the same channel it came from
func SendResult(result *SlackResult) error {

	return nil
}
