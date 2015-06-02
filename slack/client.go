package slack

import (
  "fmt"
  "errors"
  "net/http"
  "encoding/json"
	"github.com/davidhampgonsalves/slackline/conf"
)

type Response struct {
  Ok bool
  Error string
}

func PostMessage(config *conf.Conf) error {
  c := *config
  url := fmt.Sprintf("https://slack.com/api/chat.postMessage?token=%v&channel=%v&text=%v&as_user=true&parse=true&unfurl_links=false&pretty=1", c.SlackToken, c.SlackChannels[0], c.Message)
  jsonResp, err := http.Get(url)
  if err != nil {
    return errors.New("failed to post message: " + err.Error())
  }

  jsonDecoder := json.NewDecoder(jsonResp.Body)
  resp := Response{}
  err = jsonDecoder.Decode(&resp)
  if err != nil {
    return errors.New("failed to unmarshal json response: " + err.Error())
  }

  if !resp.Ok {
    return errors.New("slack error posting message: " + resp.Error)
  }

  return nil
}
