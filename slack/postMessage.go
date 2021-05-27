package slack

import (
	"github.com/slack-go/slack"
)

func (a *Api) PostMessage(channel string, text string, opts ...slack.MsgOption) error {
	opts = append(opts,
		slack.MsgOptionPostMessageParameters(slack.PostMessageParameters{Username: "SRE Hero updater", LinkNames: 1}),
	)
	_, _, err := a.api.PostMessage(channel, opts...)
	return err
}
