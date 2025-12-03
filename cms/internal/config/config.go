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
	var menuItems []*MenuItem
	/*
	 * Build menu tree
	 */
	viper.SetConfigFile("config_menu." + env + ".yaml")
	if err := viper.ReadInConfig(); err == nil {
		if err := viper.Unmarshal(&menuItems); err == nil {
			for i, menuItem := range menuItems {
				if len(menuItem.Children) > 0 {
					menuItems[i].Roles = []int{}
					for _, childMenu := range menuItem.Children {
						for _, role := range childMenu.Roles {
							hadRole := false
							for _, parentRole := range menuItem.Roles {
								if parentRole == role {
									hadRole = true
									break
								}
							}
							if hadRole {
								continue
							}
							menuItems[i].Roles = append(menuItems[i].Roles, role)
						}
					}
				}
			}
		}
	}

	return menuItems
}
