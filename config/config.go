package config

import (
	"log"

	"github.com/spf13/viper"
)

// Cfg is a global configuration object
var Cfg Config

// Config contains the server configuration
type Config struct {
	Attributes   []string `json:"attributes"`
	Base         string   `json:"base"`
	Host         string   `json:"host"`
	Port         int      `json:"port"`
	BindDN       string   `json:"bind_dn"`
	BindPassword string   `json:"bind_password"`
	UserFilter   string   `json:"user_filter"`
	SkipTLS      bool     `json:"skip_tls"`
	UseSSL       bool     `json:"use_ssl"`
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
