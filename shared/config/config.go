package config

import (
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

type Config struct {
	MetadataService MetadataService
	ApiGateway      ApiGateway
	CMS             CMS
	Web             Web
}

type Endpoint struct {
	LocalAddress  string
	PublicAddress string
}

type MetadataService struct {
	GRPC Endpoint
}

type ApiGateway struct {
	GRPC Endpoint
	HTTP Endpoint
}

type CMS struct {
	HTTP Endpoint
}

type Web struct {
	HTTP Endpoint
}

func Load(env string) (*Config, error) {
	var config Config

	_, filename, _, _ := runtime.Caller(0)
	baseDir := filepath.Dir(filename)

	viper.SetConfigFile(baseDir + "/../config." + env + ".yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	/*
	 * Override config with local config
	 */
	viper.SetConfigFile(baseDir + "/../localconfig." + env + ".yaml")

	if err := viper.ReadInConfig(); err == nil {
		viper.Unmarshal(&config)
	}

	return &config, nil
}
