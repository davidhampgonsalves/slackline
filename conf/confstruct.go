package conf

import (
  "errors"
)

type Conf struct {
	Save, Init    bool "-"
	SlackToken    string
	Message       string "-"
	SlackChannels []string
}

// Validate conf
func (c *Conf) Valid() (errs []error) {
	var conf = *c
  errs = make([]error, 0)

  if conf.Save && conf.Init {
      errs = append(errs, errors.New("Save and Init together do not make sense, you probably just want Init."))
  }
  if conf.Save {
    if len(conf.SlackChannels) == 0 && conf.SlackToken == "" {
      errs = append(errs, errors.New("Either a slack channel(-c) or token(-t) must be set to be saved."))
    }
    return
  }

  if conf.Init {
    return
  }

	if len(conf.SlackChannels) == 0 {
    errs = append(errs, errors.New("No slack channel was provided(-c) or configured."))
	}
	
  if conf.SlackToken == "" {
    errs = append(errs, errors.New("No slack token was provided(-f) or configured."))
  }
  
  return
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
