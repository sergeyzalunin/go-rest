package main

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/sergeyzalunin/go-rest/internal/app/server"
)

func main() {
	args := newArgs()
	conf := server.NewConfig()

	_, err := toml.DecodeFile(args.ConfigPath, conf)
	if err != nil {
		log.Fatal(err)
	}

	if databaseURL := os.Getenv("DATABASE_URL"); databaseURL != "" {
		conf.DatabaseURL = databaseURL
	}

	if err := server.Start(conf); err != nil {
		log.Fatal(err)
	}
}
