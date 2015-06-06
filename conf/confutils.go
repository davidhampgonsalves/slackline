package conf

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"
	"github.com/alecthomas/kingpin"
	"github.com/davidhampgonsalves/slackline/crypto"
	"github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v2"
)

func ReadFlagConfig() (conf *Conf, err error) {
	app := kingpin.New("slackline", "Fire quick status updates to slack.")
	save := app.Flag("save", "Persist current settings.").Bool()
	init := app.Flag("init", "Prompt user for settings.").Bool()
	token := app.Flag("token", "Slack auth token.").Short('t').String()
	channelsStr := app.Flag("channels", "Channel to post to.").Short('c').String()
	msgParts := app.Arg("message", "message to post to slack.").Strings()

	app.Version("slackline ver. ☃")
	app.Terminate(func(status int) {
		os.Exit(0)
	})
	kingpin.MustParse(app.Parse(os.Args[1:]))

	var channels []string
	if *channelsStr != "" {
		channels = strings.Split(*channelsStr, ",")
	}
	msg := strings.Join(*msgParts, " ")
	return &Conf{SlackChannels: channels, Save: *save, Init: *init, SlackToken: *token, Message: msg}, nil
}

func ReadYamlConf(path string) (*Conf, error) {
	path, err := homedir.Expand(path)

	// if not config file exists we load an empty config
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return new(Conf), nil
	}

	yamlData, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.New("YAML config file error " + err.Error())
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

	path, _ = homedir.Expand(path)
	err = ioutil.WriteFile(path, yamlData, 0644)
	if err != nil {
		return errors.New("failed to write YAML config: " + err.Error())
	}

	return nil
}
