// Package util handles utility functions like loading configuration
package util

import "github.com/spf13/viper"

type Config struct {
	Port       string `mapstructure:"SERVER_PORT"`
	APIKey     string `mapstructure:"STEAM_API_KEY"`
	SteamAppID string `mapstructure:"STEAM_APP_ID"`
	Env        string `mapstructure:"ENV"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath("./")
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
