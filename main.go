package main

import (
	"github.com/jtblin/go-ldap-client"
	"github.com/parsaaes/go-freeipa-example/config"
	"log"
)

func main() {
	config.Init("config.json")
	client := &ldap.LDAPClient{
		Base: config.Cfg.Base,
		Host: config.Cfg.Host,
		Port: config.Cfg.Port,
	}
	defer client.Close()

	ok, user, err := client.Authenticate(config.Cfg.User, config.Cfg.Pass)
	if err != nil {
		log.Fatalf("Error authenticating user %s: %+v", config.Cfg.User, err)
	}
	if !ok {
		log.Fatalf("Authenticating failed for user %s", config.Cfg.User)
	}
	log.Printf("User: %+v", user)
}
