package config

import (
	"github.com/spf13/viper"
)

var Config Configuration // global variable

type Configuration struct {
	Db           DBConf
	Server       ServerConf
	AccessSecret string
}

type DBConf struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

type ServerConf struct {
	Host string
	Post string
}

func LoadConfig() (err error) {
	viper.AddConfigPath("./")
	viper.SetConfigName("dev")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&Config)
	return
}
