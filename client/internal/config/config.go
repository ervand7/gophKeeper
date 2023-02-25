// Package config - client configurations.
package config

import (
	"flag"
	"log"

	"github.com/caarlos0/env/v6"
)

var servAddrFlag *string

func init() {
	servAddrFlag = flag.String("a", "", "Server address")
}

type config struct {
	ServerAddress string `env:"SERVER_ADDRESS" envDefault:":8080" json:"server_address"`
}

func GetServerAddress() string {
	var cfg config
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	flag.Parse()
	if *servAddrFlag != "" {
		cfg.ServerAddress = *servAddrFlag
	}

	return cfg.ServerAddress
}
