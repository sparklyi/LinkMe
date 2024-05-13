package config

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func InitViper() {
	configFile := pflag.String("config", "config/config.yaml", "配置文件路径")
	pflag.Parse()
	viper.SetConfigFile(*configFile)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
