package config

import (
	"shared/pkg/logger"

	"github.com/spf13/viper"
)

type Config struct {
	Port           int
	Log            logger.LogConfig
	SupportedLangs []string
	DefaultLang    string
	AllowedHosts   []string
}

func Load(env string) (*Config, error) {
	var config Config

	viper.SetConfigFile("config." + env + ".yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	/*
	 * Override config with local config
	 */
	viper.SetConfigFile("localconfig." + env + ".yaml")

	if err := viper.ReadInConfig(); err == nil {
		viper.Unmarshal(&config)
	}

	return &config, nil
}
