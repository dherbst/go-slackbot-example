package slackbot

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// Hook is the url to send information into slack for a particular domain and channel
type Hook struct {
	Url         string
	TeamDomain  string
	ChannelId   string
	ChannelName string
}

// GetHook looks up the hook url for the teamdomain
func GetHook(result *SlackResult) (*Hook, error) {

	// look up the url based on the incoming message criteria

	return &Hook{
		Url:         "https://hooks.example.com/services/TeamId/UserId/ChannelId",
		TeamDomain:  "team.example.com",
		ChannelId:   "C0001",
		ChannelName: "general",
	}, nil
}

// PostHook sends the json payload to the slack incoming hook url
func PostHook(hook *Hook, result *SlackResult) error {

	postbody, err := json.Marshal(result)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", hook.Url, bytes.NewBuffer(postbody))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
