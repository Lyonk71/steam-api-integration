// Package util handles utility functions like loading configuration
package util

import "github.com/spf13/viper"

type Config struct {
	Port           string `mapstructure:"SERVER_PORT"`
	APIKey         string `mapstructure:"STEAM_API_KEY"`
	SteamAppID     string `mapstructure:"STEAM_APP_ID"`
	Env            string `mapstructure:"ENV"`
	SteamInterface string
	SteamAPIUrl    string
}

func LoadConfig() (config Config, err error) {

	viper.SetConfigName("config") // config file name without extension
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./util/") // config file path
	viper.AutomaticEnv()           // read value ENV variable

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	if viper.Get("Env") != "PROD" {
		viper.SetDefault("SteamInterface", "ISteamMicroTxnSandbox")
	} else {
		viper.SetDefault("SteamInterface", "ISteamMicroTxn")
	}

	viper.SetDefault("SteamAPIUrl", "https://partner.steam-api.com/")

	err = viper.Unmarshal(&config)
	if err != nil {
		return
	}

	return
}
