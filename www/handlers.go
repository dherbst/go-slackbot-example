package slackbot

import (
	"net/http"
)

func init() {
	http.HandleFunc("/api/1/command", CommandHandler)
	http.HandleFunc("/", IndexHandler)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/index.html", http.StatusFound)
}
