package env

import (
	"os"

	"github.com/go-yaml/yaml"
)

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"server"`
	Logs struct {
		Directory string `yaml:"directory"`
		Filename  string `yaml:"filename"`
	} `yaml:"logs"`
}

var _cfg *Config

func GetConfig() (*Config, error) {
	if _cfg != nil {
		return _cfg, nil
	}

	f, err := os.Open("app.yaml")
	if err != nil {
		return nil, err
	}

	defer f.Close()
	var cfg Config
	dec := yaml.NewDecoder(f)
	err = dec.Decode(&cfg)
	if err != nil {
		return nil, err
	}

	_cfg = &cfg

	return &cfg, nil
}
