package conf

// Load conf from various sources and determine current settings
func Load() (*Conf, error) {
  yamlConf, err := ReadYamlConf("~/.slackline.yaml")
  if err != nil {
  	return nil, err
  }
  flagConf, err := ReadFlagConfig()
  if err != nil {
  	return nil, err
  }

  return yamlConf.Merge(flagConf), nil
}
