package config

import (
  "strings"
  "github.com/alecthomas/kingpin"
)

type FlagConfig struct {}

func (*FlagConfig) read() *Configuration {
  save := kingpin.Flag("save", "Persist current settings.").Short('s').Bool()
  init := kingpin.Flag("init", "Prompt user for settings.").Short('i').Bool()
  channelsRaw := kingpin.Flag("channels", "Channel(s) to post to(comma seperated).").Short('c').String()

  token := kingpin.Flag("token", "Slack auth token.").Short('t').String()

  kingpin.Version("0.0.1")
	kingpin.Parse()

  channels := strings.Split(*channelsRaw, ",")

  return &Configuration{channels: channels, save: *save, init: *init, token: *token}
}
