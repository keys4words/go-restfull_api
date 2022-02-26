package main

import (
	"log"

	"github.com/BurntSushi/toml"
	"github.com/keys4words/go-restfull_api/standardWebserver/internal/app/api"
)

var (
	configPath string = "configs/api.toml"
)

func init() {

}

func main() {
	log.Println("It works")
	config := api.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Println("cannot parse toml file, use default", err)
	}
	server := api.New(config)
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
