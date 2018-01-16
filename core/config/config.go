package config

import (
	"github.com/syariatifaris/gosandbox/core/config/toml"
)

var cfgData *ConfigurationData

type Config interface {
	DecodeConfig(c interface{}) error
}

func NewConfiguration() *ConfigurationData {
	var cfg ConfigurationData
	configuration := toml.NewTomlConfiguration("files/conf/app.toml")
	err := configuration.DecodeConfig(&cfg)

	if err != nil {
		return nil
	}

	return &cfg
}
