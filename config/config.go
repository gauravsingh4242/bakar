package config

import (
	"fmt"
	"github.com/gauravsingh4242/bakar/logger"
	"os"

	"github.com/spf13/viper"
)

const (
	ConfigurationFile = "BakarConfig"
	ENVPrefix         = "bakar"
)

func InitConfig() error {
	logger.Log.Info("Initializing configuration")
	viper.SetConfigType("yaml")
	viper.SetConfigFile(os.Getenv(ConfigurationFile))
	viper.SetEnvPrefix(ENVPrefix)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Failed to load config")
		return err
	}

	return nil
}
