package slackbot

import (
	_ "net/http"
	"testing"
)

func TestGetHook(t *testing.T) {
	hook, err := GetHook(nil)
	if err != nil {
		t.Fatalf("got error getting hook %v", err)
	}
	if hook.Url != "https://hooks.example.com/services/TeamId/UserId/ChannelId" {
		t.Fatalf("Did not get expected url, got %v", hook.Url)
	}
}

func TestPostHook(t *testing.T) {
	hook, err := GetHook(nil)
	if err != nil {
		t.Fatalf("got error getting hook %v", err)
	}
	result := &SlackResult{
		Text: "",
		Attachments: []*SlackAttachment{
			&SlackAttachment{ImageUrl: "http://example.com/mygif.gif"},
		},
	}

	err = PostHook(hook, result)
	if err != nil {
		t.Fatalf("Error in PostHook %v", err)
	}
}
