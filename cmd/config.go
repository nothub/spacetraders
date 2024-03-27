package main

import (
	"os"
	"path/filepath"
	"time"

	"github.com/nothub/spacetraders/internal/filez"
	"github.com/nothub/spacetraders/pkg/st"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Symbol       string           `yaml:"symbol"`
	Email        string           `yaml:"email"`
	Faction      st.FactionSymbol `yaml:"faction"`
	Token        string           `yaml:"token"`
	TokenCreated time.Time        `yaml:"token-created"`
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

func ConfigLoad() {
	ilog.Println("loading config...")

	err := filez.Touch(configPath())
	if err != nil {
		elog.Fatalln(err.Error())
	}

	data, err := os.ReadFile(configPath())
	if err != nil {
		elog.Fatalln(err.Error())
	}

	Cfg = Config{}
	err = yaml.Unmarshal(data, &Cfg)
	if err != nil {
		elog.Fatalln(err.Error())
	}
}

func ConfigSave() {
	ilog.Println("saving config...")

	data, err := yaml.Marshal(Cfg)
	if err != nil {
		elog.Fatalln(err.Error())
	}

	err = os.WriteFile(configPath(), data, 0640)
	if err != nil {
		elog.Fatalln(err.Error())
	}
}
