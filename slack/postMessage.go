package slack

import (
	"fmt"

	"github.com/slack-go/slack"
)

func (a *Api) PostMessage(channel string, text string) error {
	_, _, err := a.api.PostMessage(
		channel,
		slack.MsgOptionText(text, false),
		slack.MsgOptionPostMessageParameters(slack.PostMessageParameters{LinkNames: 1}))

	if err != nil {
		return fmt.Errorf("error posting message: %s", err)
	}
	return nil
}
