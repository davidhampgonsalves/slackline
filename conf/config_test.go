package conf

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestConfMerge(t *testing.T) {
	conf := Conf{SlackToken: "a", SlackChannels: []string{"#test"}}
	conf2 := Conf{SlackToken: "b", SlackChannels: []string{"#testing", "#testing2"}}

	merged := conf.Merge(&conf2)
	assert.Equal(t, merged.SlackToken, "b", "merged config didn't allow overriding of the slackToken")
	assert.Equal(t, len(merged.SlackChannels), 2, "merged config didn't allow overriding of slackChannels")
}

func TestFileReader(t *testing.T) {
	conf := Conf{SlackToken: "12345123451234512345", SlackChannels: []string{"#testing", "moretesting"}}
	path := "/tmp/.slackline.yml"

	err := WriteYamlConf(path, &conf)
	assert.Equal(t, err, nil, fmt.Sprintf("there was an error writing yaml config file: %v", err))

	readConf, err := ReadYamlConf(path)
	conf = *readConf

	os.Remove(path)

	assert.Equal(t, err, nil, fmt.Sprintf("there was an error reading yaml config file: %v", err))
	assert.Equal(t, conf.SlackToken, "12345123451234512345", "slack token read from config file did not match")
	assert.Equal(t, conf.SlackChannels, []string{"#testing", "moretesting"}, "slack channels read from config file did not match")
}
