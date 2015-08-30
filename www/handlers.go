package slackbot

import (
	"encoding/json"
	"net/http"
)

var version string = "1.0.0"

func init() {
	http.HandleFunc("/api/1/command", CommandHandler)
	http.HandleFunc("/version", VersionHandler)
	http.HandleFunc("/", IndexHandler)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/index.html", http.StatusFound)
}

func VersionHandler(w http.ResponseWriter, r *http.Request) {

	value := map[string]interface{}{
		"Version": version,
	}

	js, err := json.Marshal(value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
