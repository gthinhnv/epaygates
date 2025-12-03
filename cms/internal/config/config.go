package config

import (
	"shared/pkg/logger"

	"github.com/spf13/viper"
)

type Config struct {
	Port      int
	Log       logger.LogConfig
	Version   string
	MenuItems []*MenuItem
}

type MenuItem struct {
	ID        uint64
	Name      string
	Route     string
	ParentID  uint64
	Icon      string
	Level     int
	SortOrder int
	Roles     []int
	IsShow    bool

	Parent   *MenuItem
	Children []MenuItem
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

	config.MenuItems = loadMenu(env)

	return &config, nil
}

func loadMenu(env string) []*MenuItem {
	type menu struct {
		Items []*MenuItem
	}

	var m menu

	// Load config file
	viper.SetConfigFile("config_menu." + env + ".yaml")
	if err := viper.ReadInConfig(); err != nil {
		return nil
	}

	// Unmarshal YAML into struct
	if err := viper.Unmarshal(&m); err != nil {
		return nil
	}

	// Build parent roles from child roles
	for i, parent := range m.Items {
		roleSet := make(map[int]struct{})
		for _, child := range parent.Children {
			for _, role := range child.Roles {
				roleSet[role] = struct{}{}
			}
		}

		if len(roleSet) > 0 {
			// Convert map keys to slice
			m.Items[i].Roles = make([]int, 0, len(roleSet))
			for role := range roleSet {
				m.Items[i].Roles = append(m.Items[i].Roles, role)
			}
		}
	}

	return m.Items
}
