package config

import (
  "strings"
  "github.com/davidhampgonsalves/slackline/crypto"
  "github.com/alecthomas/kingpin"
)

func Load() Config { 
}

func fromFlags() Config {
  save := kingpin.Flag("save", "Persist current settings.").Short('s').Bool()
  init := kingpin.Flag("init", "Prompt user for settings.").Short('i').Bool()
  channelsRaw := kingpin.Flag("channels", "Channel(s) to post to(comma seperated).").Short('c').String()

  token := kingpin.Flag("token", "Slack auth token.").Short('t').String()

  kingpin.Version("0.0.1")
	kingpin.Parse()

  channels := strings.Split(*channelsRaw, ",")

  return Config{channels: channels, save: *save, init: *init, token: *token}
}

func fromFile() Config {
  s := crypto.Decrypt("")
}

type Config struct {
  save, init bool
  token string
  channels []string
}
