package slackbot

// Hook is the url to send information into slack for a particular domain and channel
type Hook struct {
	Url         string
	TeamDomain  string
	ChannelId   string
	ChannelName string
}

func GetHook(result *SlackResult) (*Hook, error) {

	return &Hook{
		Url:         "https://hooks.example.com/services/TeamId/UserId/ChannelId",
		TeamDomain:  "team.example.com",
		ChannelId:   "C0001",
		ChannelName: "general",
	}, nil
}

func PostHook(hook *Hook, result *SlackResult) error {

	return nil
}
