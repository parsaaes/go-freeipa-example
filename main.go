package main

import (
	"flag"
	"log"

	"github.com/jtblin/go-ldap-client"
	"github.com/parsaaes/go-freeipa-example/config"
)

func main() {
	config.Init("config.json")

	username := flag.String("username", "riemann", "LDAP username")
	password := flag.String("password", "password", "LDAP password")
	flag.Parse()

	client := &ldap.LDAPClient{
		Attributes:   config.Cfg.Attributes,
		Base:         config.Cfg.Base,
		BindDN:       config.Cfg.BindDN,
		BindPassword: config.Cfg.BindPassword,
		Host:         config.Cfg.Host,
		UserFilter:   config.Cfg.UserFilter,
		Port:         config.Cfg.Port,
		UseSSL:       config.Cfg.UseSSL,
		SkipTLS:      config.Cfg.SkipTLS,
	}
	defer client.Close()

	ok, userAttr, err := client.Authenticate(*username, *password)
	if err != nil {
		log.Fatalf("Error authenticating user %s: %+v", *username, err)
	}
	if !ok {
		log.Fatalf("Authenticating failed for user %s", *username)
	}
	log.Printf("User LDAP Attributes: %+v", userAttr)
}
