package config

import (
	"log"

	"github.com/spf13/viper"
)

// Cfg is a global configuration object
var Cfg Config

// Config contains the server configuration
type Config struct {
	Attributes   []string `mapstructure:"attributes"`
	Base         string   `mapstructure:"base"`
	Host         string   `mapstructure:"host"`
	Port         int      `mapstructure:"port"`
	BindDN       string   `mapstructure:"bind_dn"`
	BindPassword string   `mapstructure:"bind_password"`
	UserFilter   string   `mapstructure:"user_filter"`
	SkipTLS      bool     `mapstructure:"skip_tls"`
	UseSSL       bool     `mapstructure:"use_ssl"`
}

// Init initializes global configuration
func Init(path string) {
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error on reading config file: %s", err)
	}
	if err := viper.Unmarshal(&Cfg); err != nil {
		log.Fatalf("error on unmarshaling config file: %s", err)
	}
}
