package conf

type Conf struct {
	Save, Init    bool "-"
	SlackToken    string
	Message       string "-"
	SlackChannels []string
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

	if conf2.Message != "" {
		conf.Message = conf2.Message
	}

	if conf2.SlackToken != "" {
		conf.SlackToken = conf2.SlackToken
	}

	if len(conf2.SlackChannels) > 0 {
		conf.SlackChannels = conf2.SlackChannels
	}
	return &conf
}
