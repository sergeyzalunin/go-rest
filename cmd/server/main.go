package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/sergeyzalunin/go-rest/internal/app/server"
)

var (
	configString string
)

// TODO: move out this config into the separate module.
func init() {
	flag.StringVar(&configString, "config-path", "configs/config.toml", "path to config file")
}

func main() {
	flag.Parse()

	conf := server.NewConfig()

	_, err := toml.DecodeFile(configString, conf)
	if err != nil {
		log.Fatal(err)
	}

	if err := server.Start(conf); err != nil {
		log.Fatal(err)
	}
}
