package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

var config *Config

func LoadConfig() *Config {
	if config != nil {
		return config
	}

	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore and use environment variables or defaults
			fmt.Println("config.yaml file not found, using environment variables or defaults")
		} else {
			panic(fmt.Errorf("fatal error config file: %w", err))
		}
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		panic(fmt.Errorf("unable to decode into struct: %w", err))
	}

	return config
}
