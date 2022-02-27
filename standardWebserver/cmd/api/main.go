package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/keys4words/go-restfull_api/standardWebserver/internal/app/api"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "path", "configs/api.toml", "path to config file in toml format")
}

func main() {
	flag.Parse()
	log.Println("It works")
	config := api.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Println("cannot parse toml file, use default value ", err)
	}

	server := api.New(config)
	log.Fatal(server.Start())
}
