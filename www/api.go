package slackbot

import (
	"fmt"
	"net/http"
)

// Dispatch the command based on the parameter
func CommandHandler(w http.ResponseWriter, r *http.Request) {

	command := r.FormValue("command")
	text := r.FormValue("text")

	fmt.Fprintf(w, "Hello there command=%v text=%v", command, text)
}
