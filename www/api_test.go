package slackbot

import (
	"appengine/aetest"
	"testing"
)

func TestDispatchCommand(t *testing.T) {
	c, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

}
