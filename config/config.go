package config

import (
	"github.com/spf13/viper"
	"log"
)

var Cfg Config

type Config struct {
	Base string
	Host string
	Port int
	User string
	Pass string
}

func Init(path string) {
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	Cfg.Base = viper.GetString("base")
	Cfg.Host = viper.GetString("host")
	Cfg.Port = viper.GetInt("port")
	Cfg.User = viper.GetString("debug-user")
	Cfg.Pass = viper.GetString("debug-pass")
}
