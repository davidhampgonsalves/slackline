package conf

import (
  "fmt"
  "testing"
  "io/ioutil"
  "os"
  "github.com/stretchr/testify/assert"
)

func TestConfMerge(t *testing.T) {
  conf := Conf{SlackToken: "a", SlackChannels: []string{"#test"}}
  conf2 := Conf{SlackToken: "b", SlackChannels: []string{"#testing", "#testing2"}}

  merged := conf.Merge(&conf2)
  assert.Equal(t, merged.SlackToken, "b", "merged config didn't allow overriding of the slackToken")
  assert.Equal(t, len(merged.SlackChannels), 2, "merged config didn't allow overriding of slackChannels")
}

func TestFileReader(t *testing.T) {
  var data = `
  slacktoken: 12345
  slackchannels:
  - testing
  - moretesting
  `
  path := "/tmp/.slackline.yml"
  yaml := []byte(data)
  ioutil.WriteFile(path, yaml, 0644)

  conf, err := ReadYamlConf(path)
  os.Remove(path)

  assert.Equal(t, err, nil, fmt.Sprintf("there was an error reading config file: %v", err))
  assert.Equal(t, conf.SlackToken, "12345", "slack token read from config file did not match")
  assert.Equal(t, conf.SlackChannels, []string{"testing", "moretesting"}, "slack channels read from config file did not match")
}
