package config

import (
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	once            sync.Once
	config_Instance *Config
	env             *Enviroment
)

func GetConfig() *Config {
	once.Do(func() {
		viper.SetConfigFile(".env")
		if err := viper.ReadInConfig(); err != nil {
			logrus.Fatalf("Failed to load the configuration file from disk and key/value stores, searching in one of the defined paths : %s\n", err)
		}

		if err := viper.ReadInConfig(); err != nil {
			logrus.Fatalf("Failed to load enviroment : %s\n", err)
		}

		if err := viper.Unmarshal(&env); err != nil {
			logrus.Fatalf("Failed to decode the enviroment : %s\n", err)
		}

		if err := validator.New().Struct(env); err != nil {
			logrus.Fatalf("Failed to validate the enviroment : %s\n", err)
		}
		// fmt.Printf("%+v\n", env)
		config_Instance = &Config{Env: env}
	})
	// fmt.Println(config_Instance)
	return config_Instance
}
