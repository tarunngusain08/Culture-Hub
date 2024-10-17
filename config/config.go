package config

import (
	"fmt"

	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/tarunngusain08/culturehub/pkg/log"
)

var logger = log.New("config")

// Start sets up the configuration and logger
// need to call this in main
func Startup() error {
	file := fmt.Sprintf("%s_%s", GetEnv(), "config")
	logger.Info("Reading config file", zap.String("file", file))

	viper.SetConfigName(file)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(GetAppPath() + "config")
	return viper.ReadInConfig()
}

// GetBool gets the value of the given key as a bool
func GetBool(key string) bool {
	return viper.GetBool(key)
}

// GetString returns the value of the given key as a string
func GetString(key string) string {
	return viper.GetString(key)
}

// GetInt returns the value of the given key as an int
func GetInt(key string) int {
	return viper.GetInt(key)
}

// SetBool sets the value of the given key to the given bool
func Set(key string, val any) {
	viper.Set(key, val)
}
