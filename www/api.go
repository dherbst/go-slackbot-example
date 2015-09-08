package slackbot

import (
	"appengine"
	"errors"
	"fmt"
	"net/http"
	"strings"
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

// SlackResult holds the result of processing the command.  json encoding is the `payload`
// message to a slack incoming hook integration.
type SlackResult struct {
	Command      *SlackCommand      `json:"-"` // ignored
	IsTextResult bool               `json:"-"`
	Text         string             `json:"text"`
	Username     string             `json:"username,omitempty"`
	IconUrl      string             `json:"icon_url,omitempty"`
	IconEmoji    string             `json:"icon_emoji,omitempty"`
	Channel      string             `json:"channel,omitempty"`
	Attachments  []*SlackAttachment `json:"attachments,omitempty"`
}

// SlackAttachment is a message attachment
type SlackAttachment struct {
	ImageUrl string `json:"image_url,omitempty"`
	ThumbUrl string `json:"thumb_url,omitempty"`
	Text     string `json:"text,omitempty"`
	Fallback string `json:"fallback,omitempty"`
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

	result, err := ProcessCommand(command) // HL
	if err != nil {
		http.Error(w, fmt.Sprintf("Error processing result %v", err), http.StatusInternalServerError)
		return
	}
	if result.IsTextResult { // HL
		fmt.Fprintf(w, "%v", result.Text)
	} else {
		c := appengine.NewContext(r)
		err = SendResult(c, result) // HL
		if err != nil {
			http.Error(w, fmt.Sprintf("Error processing result %v", err), http.StatusInternalServerError)
			return
		}
		http.Error(w, "", http.StatusNoContent) // return a 204
	}
}

// Determine what do to with the different text commands
func ProcessCommand(cmd *SlackCommand) (*SlackResult, error) {
	var result *SlackResult
	var err error
	parts := strings.Split(cmd.Text, " ")
	commandWord := parts[0]
	switch commandWord { // HL
	case "hello":
		result, err = HelloCommand(cmd)
	case "gif":
		result, err = GifCommand(cmd)
	default:
		result, err = UnknownCommand(cmd)
	}
	return result, err
}

// HelloCommand says hi back to the user that said hello
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
func SendResult(c appengine.Context, result *SlackResult) error {
	if result.IsTextResult {
		return errors.New("Sending text result to non-text send")
	}

	hook, err := GetHook(result)
	if err != nil {
		return err
	}

	err = PostHook(c, hook, result)

	return err
}

// GifCommand looks up the gif in the datastore and returns the url to it
func GifCommand(cmd *SlackCommand) (*SlackResult, error) {

	//a := &SlackAttachment{ImageUrl: "http://dramafeverslack.appspot.com/gif/heirs7_1.gif"}
	//attachments := make([]*SlackAttachment, 1)
	//attachments[0] = a
	result := &SlackResult{
		IsTextResult: false,
		Username:     "DramaFever",
		IconUrl:      "http://slack.dramafever.com/gif/df-flame.png",
		Text:         cmd.Text,
		Attachments: []*SlackAttachment{
			&SlackAttachment{
				ImageUrl: "http://dramafeverslack.appspot.com/gif/heirs7_1.gif",
				ThumbUrl: "http://dramafeverslack.appspot.com/gif/heirs7_1.gif",
			},
		},
		Command: cmd,
	}
	return result, nil
}
