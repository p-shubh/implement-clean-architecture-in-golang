package config

import (
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	once            sync.Once
	config_Instance *Config
)

func GetConfig() *Config {
	once.Do(func() {
		var env Enviroment
		viper.SetConfigFile(".env")
		viper.ReadInConfig()

		if err := viper.ReadInConfig(); err != nil {
			logrus.Errorf("Failed to load enviroment : %s", err)
		}

		if err := viper.Unmarshal(&env); err != nil {
			logrus.Errorf("Failed to decode the enviroment : %s", err)
		}
	})
	return config_Instance
}
