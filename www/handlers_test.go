package slackbot

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestVersionHandler(t *testing.T) {
	request, err := http.NewRequest("GET", "/api/1/version", nil)
	if err != nil {
		t.FailNow()
	}
	response := httptest.NewRecorder()

	VersionHandler(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("Status from GET did not return 200, instead returned %v", response.Code)
		t.FailNow()
	}

	versionBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.FailNow()
	}
	var obj map[string]interface{}
	err = json.Unmarshal(versionBody, &obj)
	if err != nil {
		t.Fatalf("Could not unmarshal version body %v", versionBody)
		t.FailNow()
	}
	if obj["Version"] != "1.0.0" {
		t.Fatalf("Version=%v", obj["Version"])
		t.FailNow()
	}
}
