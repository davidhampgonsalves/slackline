package conf

type Conf struct {
	Save, Init          bool "-"
	SlackToken, Message string
	SlackChannels       []string
}

// Validate conf
func (c *Conf) Valid() bool {
	var conf = *c
	if len(conf.SlackChannels) == 0 {
		return false
	}
	return conf.SlackToken != ""
}

// Merge two confs(without mutating) taking any existing vales from
// passed in Conf
func (conf Conf) Merge(c2 *Conf) *Conf {
	conf2 := *c2
	conf.Save = conf2.Save
	conf.Init = conf2.Init

	if conf2.SlackToken != "" {
		conf.SlackToken = conf2.SlackToken
	}
	if conf2.SlackChannels != nil {
		conf.SlackChannels = conf2.SlackChannels
	}
	return &conf
}
