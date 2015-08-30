package slackbot

import (
	"appengine/aetest"
	"net/http"
	"strings"
	"testing"
)

func TestDispatchCommand(t *testing.T) {
	c, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

}

// Pass a request with the form values in and check they are placed in the SlackCommand
func TestUnMarshalCommand(t *testing.T) {
	req, _ := http.NewRequest(
		"POST", "/api/1/command",
		strings.NewReader("token=tokenhere&team_id=T0001&team_domain=example.com&channel_id=C123&channel_name=test&user_id=U1234&user_name=tron&command=/fight&text=for_the_users"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	c, err := UnMarshalCommand(req)
	if err != nil {
		t.Fatalf("Got err from UnMarshalCommand=%v", err)
	}
	if c.Token != "tokenhere" {
		t.Fatalf("Did not get expected token, got %v", c.Token)
	}

}
