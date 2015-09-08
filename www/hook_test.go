package slackbot

import (
	"appengine/aetest"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
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

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("Could not read post body %v", err)
		}
		defer r.Body.Close()
		bodystr := string(body)
		t.Logf("body=%v\n", bodystr)
		if !strings.Contains(bodystr, "\"image_url\":\"http://example.com/mygif.gif\"") {
			t.Fatalf("Could not find gif")
		}

		fmt.Fprintf(w, "OK")
	}))
	defer ts.Close()

	hook, err := GetHook(nil)
	if err != nil {
		t.Fatalf("got error getting hook %v", err)
	}
	hook.Url = ts.URL

	result := &SlackResult{
		Text: "",
		Attachments: []*SlackAttachment{
			&SlackAttachment{ImageUrl: "http://example.com/mygif.gif"},
		},
	}

	c, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatalf("Error getting context %v\n", err)
	}
	err = PostHook(c, hook, result)
	if err != nil {
		t.Fatalf("Error in PostHook %v", err)
	}
}
