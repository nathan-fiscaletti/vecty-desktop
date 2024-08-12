package config

import (
	"embed"

	"gopkg.in/yaml.v2"
)

//go:embed config.yaml
var cfgFileSystem embed.FS
var cfg *config

type config struct {
	Port int `yaml:"port"`
}

func GetConfig() (*config, error) {
	if cfg != nil {
		return cfg, nil
	}

	data, err := cfgFileSystem.ReadFile("config.yaml")
	if err != nil {
		return nil, err
	}

	cfg = &config{}
	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
