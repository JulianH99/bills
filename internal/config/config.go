package config

import (
	"github.com/adrg/xdg"
	"github.com/spf13/viper"
)

type config struct {
	IconPath        string `yaml:"icon_path" mapstructure:"icon_path"`
	DatabaseCreated bool   `yaml:"database_created" mapstructure:"database_created"`
}

func Set(cfg config) error {
	viper.Set("config", cfg)
	return viper.WriteConfig()
}

func InitializeConfig() error {
	viper.SetConfigName("gobills")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(xdg.ConfigHome)

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	return nil
}

func GetConfig() (*config, error) {
	var cfg config
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.UnmarshalKey("config", &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
