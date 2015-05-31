package config

type Configuration struct {
  save, init bool
  token string
  channels []string
}

func (cPtr *Configuration) isValid() bool {
  var c = *cPtr
  if len(c.channels) == 0 {
    return false
  }
  if c.token == "" {
    return false
  }
  return true
}

func (c Configuration) merge(c2Ptr *Configuration) *Configuration {
  c2 := *c2Ptr
  c.save = c2.save
  c.init = c2.init

  if c2.token != "" {
    c.token = c2.token
  }
  if c2.channels != nil {
    c.channels = c2.channels
  }
  return &c
}

type ConfigReader interface {
  read() *Configuration
}

type ConfigWriter interface {
  write(*Configuration)
}

func Load() *Configuration {
  configReaders := []ConfigReader{new(FileConfig), new(FlagConfig)}
  return configReaders[0].read()
}


