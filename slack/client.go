package slack

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/davidhampgonsalves/slackline/conf"
	"net/http"
	"net/url"
)

func PostMessage(config *conf.Conf) error {
	c := *config
	msg := formatMessage(config)
	url := fmt.Sprintf("https://slack.com/api/chat.postMessage?token=%v&channel=%v&text=%v&as_user=true&link_names=1&unfurl_links=false", c.SlackToken, url.QueryEscape(c.SlackChannels[0]), msg)
	return makeSlackRequest("post message", url)
}

func JoinChannel(config *conf.Conf) error {
	c := *config
	url := fmt.Sprintf("https://slack.com/api/channels.join?token=%v&name=%v", c.SlackToken, url.QueryEscape(c.SlackChannels[0]))
	return makeSlackRequest("join channel", url)
}

func formatMessage(config *conf.Conf) string {
	msg := config.Message
	return url.QueryEscape(msg)
}

type Response struct {
	Ok    bool
	Error string
}

func makeSlackRequest(action string, url string) error {
	jsonResp, err := http.Get(url)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to %v: %v", action, err.Error()))
	}

	jsonDecoder := json.NewDecoder(jsonResp.Body)
	resp := Response{}
	err = jsonDecoder.Decode(&resp)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to unmarshal %v response: %v", action, err.Error()))
	}

	if !resp.Ok {
		return errors.New(fmt.Sprintf("slack error %v: %v", action, resp.Error))
	}

	return nil
}
