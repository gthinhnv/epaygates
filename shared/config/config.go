package config

import "github.com/spf13/viper"

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
