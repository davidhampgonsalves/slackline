package conf

import (
  "fmt"
  "github.com/bgentry/speakeasy"
)
// Load conf from various sources and determine current settings
func Load() (*Conf, error) {
	yamlConf, err := ReadYamlConf("~/.slackline.yaml")
	if err != nil {
		return nil, err
	}

	flagConf, err := ReadFlagConfig()
	if err != nil {
		return nil, err
	}

	return yamlConf.Merge(flagConf), nil
}

func Save(c *Conf) error {
	err := WriteYamlConf("~/.slackline.yaml", c)
	return err
}

func Init(c *Conf) error {
  conf := *c

  fmt.Println("Enter your Slack Web API token(or enter to skip), you can find it at: ") 
  token, _ := speakeasy.Ask("")
  if token != ""  {
    conf.SlackToken = token
  }

  fmt.Println("Enter the default channel you want to post to.") 
  var channel string
  fmt.Scanf("%s", &channel)
  if channel != "" {
    conf.SlackChannels = []string{channel}
  }

  return Save(&conf)
}
