package config

import (
	"github.com/spf13/viper"
)

var v = viper.GetViper()

func Boot() {
	initViper()
	initApp()
	initDatabase()
}
func initViper() {
	viper.SetConfigName(".env") // allow directly reading from .env file
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("/")
	viper.AllowEmptyEnv(true)
	viper.AutomaticEnv()
	_ = viper.ReadInConfig()
}
