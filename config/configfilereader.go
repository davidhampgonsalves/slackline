package config

import (
  "github.com/davidhampgonsalves/slackline/crypto"
)

type FileConfig struct {}

func (*FileConfig) read() *Configuration {
  s := crypto.Decrypt("")
  return &Configuration{token: s}
}
