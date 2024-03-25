package main

import (
	"os"
	"path/filepath"

	"github.com/nothub/spacetraders/internal/files"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Token string `yaml:"token"`
}

var Cfg Config

var ConfigPath = ""

func configPath() (path string) {
	if ConfigPath != "" {
		return ConfigPath
	}
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err.Error())
	}
	ConfigPath = filepath.Join(home, ".config", "spacetraders", "config.yaml")
	return ConfigPath
}

func ConfigLoad() error {
	err := files.Touch(configPath())
	if err != nil {
		return err
	}

	data, err := os.ReadFile(configPath())
	if err != nil {
		return err
	}

	Cfg = Config{}
	err = yaml.Unmarshal(data, &Cfg)
	if err != nil {
		return err
	}

	return nil
}

func ConfigSave() error {
	data, err := yaml.Marshal(Cfg)
	if err != nil {
		return err
	}

	err = os.WriteFile(configPath(), data, 0640)
	if err != nil {
		return err
	}

	return nil
}
