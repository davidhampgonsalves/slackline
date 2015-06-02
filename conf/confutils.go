package conf

import (
	"errors"
	"github.com/alecthomas/kingpin"
	"github.com/davidhampgonsalves/slackline/crypto"
	"github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"
)

func ReadFlagConfig() (*Conf, error) {
	save := kingpin.Flag("save", "Persist current settings.").Short('s').Bool()
	init := kingpin.Flag("init", "Prompt user for settings.").Short('i').Bool()
	token := kingpin.Flag("token", "Slack auth token.").Short('t').String()
	channelsStr := kingpin.Flag("channels", "Channel(s) to post to(comma seperated).").Short('c').String()
	msg := kingpin.Arg("message", "message to post to slack.").Required().String()

	kingpin.Version("0.0.1")
	kingpin.Parse()

	channels := strings.Split(*channelsStr, ",")

	return &Conf{SlackChannels: channels, Save: *save, Init: *init, SlackToken: *token, Message: *msg}, nil
}

func ReadYamlConf(path string) (*Conf, error) {
	path, err := homedir.Expand(path)

	// if not config file exists we load an empty config
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return new(Conf), nil
	}

	yamlData, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.New("YAML onfig file error " + err.Error())
	}

	conf := Conf{}
	err = yaml.Unmarshal(yamlData, &conf)
	if err != nil {
		return nil, errors.New("YAML config file unmarshal error " + err.Error())
	}

	if conf.SlackToken != "" {
		conf.SlackToken, err = crypto.Decrypt(conf.SlackToken)
		if err != nil {
			return nil, errors.New("Slack token from YAML was corrupt " + err.Error())
		}
	}

	return &conf, nil
}

func WriteYamlConf(path string, c *Conf) error {
	conf := *c
	if conf.SlackToken != "" {
		conf.SlackToken = crypto.Encrypt(conf.SlackToken)
	}

	yamlData, err := yaml.Marshal(conf)
	if err != nil {
		return errors.New("failed to marshal config to YAML: " + err.Error())
	}

	err = ioutil.WriteFile(path, yamlData, 0644)
	if err != nil {
		return errors.New("failed to write YAML config: " + err.Error())
	}

	return nil
}
